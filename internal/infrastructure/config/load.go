package config

import (
	"github.com/spf13/viper"
)

var GlobalConfig *Config

func LoadConfig(path string) error {
	viper.SetConfigFile(path)
	if err := viper.ReadInConfig(); err != nil {
		return err
	}
	config := &Config{}
	if err := viper.Unmarshal(config); err != nil {
		return err
	}
	GlobalConfig = config
	return nil
}
