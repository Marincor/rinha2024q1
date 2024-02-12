package config

import (
	"log"
	"path/filepath"
	"runtime"

	"github.com/spf13/viper"
)

type Config struct {
	DBString  string `json:"database_url"`
	DBLogMode int    `json:"db_log_mode"`
}

func New() *Config {
	return setupLocal()
}

func setupLocal() *Config {
	var config *Config

	_, file, _, _ := runtime.Caller(0) //nolint: dogsled

	viper.SetConfigName("config.yaml")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(filepath.Join(filepath.Dir(file), "../"))

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error while reading config file, %s", err)
	}

	err := viper.Unmarshal(&config)
	if err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
	}
	return config
}
