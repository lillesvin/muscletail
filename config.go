package main

import (
	"github.com/BurntSushi/toml"
	log "github.com/sirupsen/logrus"
)

// Config holds the parsed config
type Config struct {
	Watches []Watch `toml:"watch"`
}

// ReadConfig accepts a path to a TOML file containing the application configuration
func ReadConfig(path string) *Config {
	cfg := &Config{}
	if _, err := toml.DecodeFile(path, cfg); err != nil {
		log.Fatalf("Failed to decode config file: %s - %s", path, err)
	}
	return cfg
}
