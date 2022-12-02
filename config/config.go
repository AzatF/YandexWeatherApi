package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"sync"
)

type Config struct {
	Token     string  `env:"token"`
	Latitude  float64 `env:"latitude"`
	Longitude float64 `env:"longitude"`
}

var (
	instance *Config
	once     sync.Once
)

func GetConfig(path string) *Config {

	once.Do(func() {
		log.Printf("Read config from: %s", path)

		instance = &Config{}

		if err := cleanenv.ReadConfig(path, instance); err != nil {
			log.Fatal(err)
		}
	})
	return instance
}
