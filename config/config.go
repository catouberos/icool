package config

import (
	"flag"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

type Config struct {
	// connection
	Host string `mapstructure:"host"`

	// params
	Room string `mapstructure:"room"`
}

func Load() (*Config, error) {
	setDefault()
	bindFlags()

	config := &Config{}

	if err := viper.Unmarshal(config); err != nil {
		return nil, err
	}

	return config, nil
}

func bindFlags() {
	flag.String("host", "", "icool management address")
	flag.String("room", "", "room")

	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	pflag.Parse()
	viper.BindPFlags(pflag.CommandLine)
}

func setDefault() {
	viper.SetDefault("host", "192.168.11.150:3000")
	viper.SetDefault("room", "41")
}
