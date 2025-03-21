package config

import (
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type GRPCServer struct {
	IP   string `yaml:"ip"`
	Port string `yaml:"port"`
}

type Config struct {
	StoragePath string     `yaml:"storage_path"`
	DBEngine    string     `yaml:"db_engine"`
	LogLevel    string     `yaml:"log_level"`
	LogOutput   string     `yaml:"log_output"`
	LogFilePath string     `yaml:"log_file_path"`
	Server      GRPCServer `yaml:"server"`
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
