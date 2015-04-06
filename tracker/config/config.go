package config

import (
	"code.google.com/p/gcfg"
	"log"
)

type Config struct {
	Server struct {
		BindAddress string
		Port        string
	}
	Database struct {
		Host     string
		Port     string
		User     string
		Password string
	}
}

func GetConfig() Config {
	var cfg Config
	err := gcfg.ReadFileInto(&cfg, "config.ini")
	if err != nil {
		log.Fatal("Configuration error: " + err.Error())
	}
	if cfg.Server.BindAddress == "" {
		cfg.Server.BindAddress = "[::]"
	}
	if cfg.Server.Port == "" {
		cfg.Server.Port = "8080"
	}
	return cfg
}
