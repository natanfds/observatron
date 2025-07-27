package utils

import (
	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

type envStruct struct {
	ApiPort string `env:"API_PORT"`
}

func (e *envStruct) Load() error {
	err := godotenv.Load()
	if err != nil {
		return err
	}

	cfg, err := env.ParseAs[envStruct]()
	if err != nil {
		return err
	}
	ENV = cfg
	ENV.ApiPort = ":" + cfg.ApiPort
	return nil
}

var ENV envStruct
