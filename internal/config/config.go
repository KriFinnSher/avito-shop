package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	Server struct {
		Host string `mapstructure:"host"`
		Port string `mapstructure:"port"`
	}
	Db struct {
		Host string `mapstructure:"host"`
		Port string `mapstructure:"port"`
		User string `mapstructure:"user"`
		Pass string `mapstructure:"pass"`
		Name string `mapstructure:"name"`
	}
	JWT struct {
		SecretKey string `mapstructure:"secret_key"`
	}
}

var AppConfig Config

func SetUpConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./internal/config/")

	if err := viper.ReadInConfig(); err != nil {
		// TODO: add wrap on this error
	}

	if err := viper.Unmarshal(&AppConfig); err != nil {
		// TODO: add wrap on this error
	}
}
