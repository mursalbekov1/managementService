package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"os"
)

type Config struct {
	env        string     `yaml:"env"`
	HttpServer HttpServer `yaml:"http-server"`
}

type HttpServer struct {
	Port string `yaml:"port"`
	Host string `yaml:"host"`
}

func MustLoad() *Config {
	configPath := os.Getenv("CONFIG_PATH")

	if configPath == "" {
		configPath = "./config/config.yaml"
	}

	var cfg Config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		panic(err)
	}

	return &cfg
}
