package config

import (
	"log"

	"github.com/spf13/viper"
)

var config *viper.Viper

func Init(env string) {

	var err error
	config = viper.New()
	config.SetConfigType("yml")
	config.SetConfigName(env)
	config.AddConfigPath("../config/")
	config.AddConfigPath("config/")

	err = config.ReadInConfig()

	if err != nil {
		log.Fatalf("Error on parsing configuration file, %s", err)
	}
}

func GetConfig() *viper.Viper {
	return config
}
