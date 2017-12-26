package beater

import (
	"fmt"
	"strings"
	"time"

	"github.com/elastic/beats/libbeat/beat"
	"github.com/elastic/beats/libbeat/common"
	"github.com/elastic/beats/libbeat/logp"

	"github.com/benben/serialbeat/config"

	"github.com/tarm/serial"
)

type Serialbeat struct {
	done         chan struct{}
	config       config.Config
	client       beat.Client
	serialConfig *serial.Config
}

// Creates beater
func New(b *beat.Beat, cfg *common.Config) (beat.Beater, error) {
	config := config.DefaultConfig
	if err := cfg.Unpack(&config); err != nil {
		return nil, fmt.Errorf("Error reading config file: %v", err)
	}

	bt := &Serialbeat{
		done:         make(chan struct{}),
		config:       config,
		serialConfig: &serial.Config{Name: config.Device, Baud: config.Baud},
	}
	return bt, nil
}

func (bt *Serialbeat) Run(b *beat.Beat) error {
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
			_, err = serial.Write([]byte(bt.config.Init[i] + "\n"))

			if err != nil {
				return err
			}
		}
	}

	serialDataReceived := make(chan bool, 1)
	go func() {
		for {
			var str string
			// read from serial as long as we didn't receive something already
			// or it didn't end with \n
			for strings.Count(str, "") <= 1 || !(strings.Contains(str, "\n")) {
				buf := make([]byte, 128)
				read, _ := serial.Read(buf)
				str += string(buf[:read])
			}

			event := beat.Event{
				Timestamp: time.Now(),
				Fields: common.MapStr{
					"type": b.Info.Name,
					"data": str,
				},
			}

			bt.client.Publish(event)
			logp.Info("Event sent")
			serialDataReceived <- true
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

func (bt *Serialbeat) Stop() {
	bt.client.Close()
	close(bt.done)
}
