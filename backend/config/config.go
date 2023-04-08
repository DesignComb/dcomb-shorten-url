package config

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var Val Config

type Config struct {
	Mode string `structure:"MODE"`
	Port string `structure:"PORT"`

	DbHost string `structure:"DB_HOST"`
	DbUser string `structure:"DB_USER"`
	DbPwd string `structure:"DB_PWD"`
	DbName string `structure:"DB_NAME"`

	GoogleSecretKey string `structure:"GOOGLE_SECRET_KEY"`
	GoogleClientID  string `structure:"GOOGLE_CLIENT_ID"`

	JWTTokenLife  int `structure:"JWT_TOKEN_LIFE"`
	JWTSecret  string `structure:"JWT_TOKEN_LIFE"`
}

func Init() {
	viper.SetConfigName("config.yaml")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		log.Panicf("Fatal error config file: %v ", err)
	}
	if err := viper.Unmarshal(&Val); err != nil {
		log.Panicf("unable to decode into struct, %v", err)
	}

	log.WithFields(log.Fields{
		"val": Val,
	}).Info("config loaded")
}
