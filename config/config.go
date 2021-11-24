package config

import (
	"log"

	"github.com/spf13/viper"
)

var (
	API_PORT    = ViperEnvVariable("API_PORT")
	DB_HOST     = ViperEnvVariable("DB_HOST")
	DB_USER     = ViperEnvVariable("DB_USER")
	DB_PASSWORD = ViperEnvVariable("DB_PASSWORD")
	DB_NAME     = ViperEnvVariable("DB_NAME")
	DB_PORT     = ViperEnvVariable("DB_PORT")
)

func ViperEnvVariable(key string) string {
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error while reading config file %s", err)
	}
	value, ok := viper.Get(key).(string)
	if !ok {
		log.Fatalf("Invalid type assertion")
	}
	return value
}
