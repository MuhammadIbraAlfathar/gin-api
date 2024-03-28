package config

import (
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	PORT        string
	DB_USERNAME string
	DB_PASSWORD string
	DB_URL      string
	DB_DATABASE string
}

var ENV *Config

func LoadConfig() {
	viper.AddConfigPath(".")
	viper.SetConfigName(".env")
	viper.SetConfigType("env")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(err)
	}

	if err := viper.Unmarshal(&ENV); err != nil {
		log.Fatal(err)
	}
}
