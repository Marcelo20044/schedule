package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"os"
	"time"
)

type Config struct {
	Env    string `yaml:"env" env-default:"local"`
	Server `yaml:"http_server"`
	Db     `yaml:"db"`
	Kafka  `yaml:"kafka"`
}

type Server struct {
	Host        string        `yaml:"host" env-default:"localhost"`
	Port        string        `yaml:"port" env-default:"9090"`
	Timeout     time.Duration `yaml:"timeout" env-default:"4s"`
	IdleTimeout time.Duration `yaml:"idle_timeout" env-default:"60s"`
}

type Db struct {
	Driver   string `yaml:"driver" env-default:"postgres"`
	DbName   string `yaml:"name" env-default:"postgres"`
	Port     string `yaml:"port" env-default:"5432"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}

type Kafka struct {
	Host        string `yaml:"host" env-default:"localhost"`
	Port        string `yaml:"port" env-default:"9092"`
	GroupsTopic string `yaml:"groups_topic" env-default:"groups"`
}

func GetConfig() *Config {
	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		log.Fatal("CONFIG_PATH is not set")
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("config file does not exist: %s", configPath)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		log.Fatalf("cannot read config: %s", err)
	}

	return &cfg
}
