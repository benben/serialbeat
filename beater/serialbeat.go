package beater

import (
	"bufio"
	"fmt"
	"time"

	"github.com/elastic/beats/v7/libbeat/beat"
	"github.com/elastic/beats/v7/libbeat/common"
	"github.com/elastic/beats/v7/libbeat/logp"

	"github.com/suda/serialbeat/config"

	"github.com/tarm/serial"
)

// serialbeat configuration.
type serialbeat struct {
	done         chan struct{}
	config       config.Config
	client       beat.Client
	serialConfig *serial.Config
}

// New creates an instance of serialbeat.
func New(b *beat.Beat, cfg *common.Config) (beat.Beater, error) {
	c := config.DefaultConfig
	if err := cfg.Unpack(&c); err != nil {
		return nil, fmt.Errorf("Error reading config file: %v", err)
	}

	bt := &serialbeat{
		done:         make(chan struct{}),
		config:       c,
		serialConfig: &serial.Config{Name: c.Device, Baud: c.Baud},
	}
	return bt, nil
}

// Run starts serialbeat.
func (bt *serialbeat) Run(b *beat.Beat) error {
	logp.Info("serialbeat is running! Hit CTRL-C to stop it.")

	var err error
	bt.client, err = b.Publisher.Connect()
	if err != nil {
		return err
	}

	serial, err := serial.OpenPort(bt.serialConfig)
	if err != nil {
		return err
	}

	if len(bt.config.Init) > 0 {
		for i := range bt.config.Init {
			_, err = serial.Write([]byte(bt.config.Init[i] + bt.config.Delimiter))

			if err != nil {
				return err
			}
		}
	}

	serialDataReceived := make(chan bool, 1)
	go func() {
		for {
			scanner := bufio.NewScanner(serial)
			for scanner.Scan() {
				if scanner.Text() != "" {
					event := beat.Event{
						Timestamp: time.Now(),
						Fields: common.MapStr{
							"type":   b.Info.Name,
							"data":   scanner.Text(),
							"device": bt.serialConfig.Name,
						},
					}

					bt.client.Publish(event)
					logp.Info("Event sent")
					serialDataReceived <- true
				}
			}
			if scanner.Err() != nil {
				_ = fmt.Errorf("Error reading serial: %v", err)
			}
		}
	}()

	for {
		select {
		case <-bt.done:
			return nil
		case <-serialDataReceived:
		}
	}
}

// Stop stops serialbeat.
func (bt *serialbeat) Stop() {
	bt.client.Close()
	close(bt.done)
}
