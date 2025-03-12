package config

import (
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type WebSocketServer struct {
	IP   string `yaml:"ip"`
	Port string `yaml:"port"`
}

type GRPCClient struct {
	IP   string `yaml:"ip"`
	Port string `yaml:"port"`
}

type Config struct {
	LogLevel    string          `yaml:"log_level"`
	LogOutput   string          `yaml:"log_output"`
	LogFilePath string          `yaml:"log_file_path"`
	Server      WebSocketServer `yaml:"server"`
	Client      GRPCClient      `yaml:"client"`
}

func MustLoad() *Config {
	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		panic("Config path must be specified")
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		panic(err.Error())
	}

	var cfg Config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		panic(err.Error())
	}

	return &cfg
}
