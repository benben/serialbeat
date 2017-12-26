// Config is put into a different package to prevent cyclic imports in case
// it is needed in several locations

package config

type Config struct {
	Device string   `config:"device"`
	Baud   int      `config:"baud"`
	Init   []string `config:"init"`
}

var DefaultConfig = Config{
	Device: "/dev/ttyACM0",
	Baud:   38400,
}
