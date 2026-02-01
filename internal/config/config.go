package config

import (
	"flag"
	"fmt"
)

type Config struct {
	FilePath string
}

func NewConfig() *Config {
	cfg := &Config{}

	flag.StringVar(&cfg.FilePath, "file", "", "путь к файлу с эндпоинтами (обязательно)")

	flag.Parse()

	return cfg
}

func (c *Config) Validate() error {
	if c.FilePath == "" {
		return fmt.Errorf("флаг -file является обязательным")
	}
	return nil
}