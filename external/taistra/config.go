package taistra

import (
	"os"
)

type Config struct {
    AmqpURL string
    TaistraURL string
}

func NewConfig() (*Config, error) {
  config := &Config{
    AmqpURL: os.Getenv("AMQP_URL"),
    TaistraURL: os.Getenv("TAISTRA_URL"),
  }

  return config, nil
}
