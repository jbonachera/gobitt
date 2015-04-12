package config

import (
	"code.google.com/p/gcfg"
	"log"
	"time"
)

type Config struct {
	Server struct {
		BindAddress    string
		Port           string
		DatabasePlugin string
		MaxPeerAge     time.Duration
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
	if cfg.Server.DatabasePlugin == "" {
		cfg.Server.DatabasePlugin = "memory"
	}
	if cfg.Server.MaxPeerAge <= 0 {
		cfg.Server.MaxPeerAge = 3600
	}
	return cfg
}
