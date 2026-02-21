package main

import (
	"github.com/spf13/viper"
)

func LoadConfig() Config {
	viper.SetDefault("port", 8080)
	viper.SetDefault("secret", "my-secret")

	return Config{
		Port:   viper.GetInt("port"),
		Secret: viper.GetString("secret"),
	}
}