package config

import (
	"github.com/spf13/viper"
	"log"
)

// Config is used to store data from config.yaml file in convenient struct
type Config struct {
	Server struct {
		Host string `mapstructure:"host"`
		Port string `mapstructure:"port"`
	}
	DB struct {
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

// SetUpConfig is used to set up global var for storing config data
func SetUpConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./internal/config/")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("failed to read config")
	}

	if err := viper.Unmarshal(&AppConfig); err != nil {
		log.Fatalf("failed to unmarshal config")
	}
}
