package config

import (
	"errors"

	"github.com/spf13/viper"
)

type Config struct {
	// Server
	ServerHost  string `mapstructure:"APP_HOST"`
	ServerPort  string `mapstructure:"APP_PORT"`
	Environment string `mapstructure:"ENVIRONMENT"`

	// Database
	DBHost         string `mapstructure:"DB_HOST"`
	DBPort         string `mapstructure:"DB_PORT"`
	DBTimeZone     string `mapstructure:"DB_TIMEZONE"`
	DBName         string `mapstructure:"DB_DATABASE"`
	DBUserName     string `mapstructure:"DB_USER"`
	DBUserPassword string `mapstructure:"DB_PASSWORD"`
}

var configInstance *Config

func Load() (err error) {
	viper.AddConfigPath(".")
	viper.SetConfigType("env")
	viper.SetConfigName("app")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&configInstance)
	return
}

func GetConfig() (Config, error) {
	if configInstance == nil {
		return Config{}, errors.New("config not loaded")
	}
	return *configInstance, nil
}
