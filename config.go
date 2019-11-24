package main

import (
	"github.com/BurntSushi/toml"
	log "github.com/sirupsen/logrus"
)

type Comparison string

const (
	GreaterThan Comparison = "GreaterThan"
	LessThan    Comparison = "LessThan"
)

type Config struct {
	Watches []Watch `toml:"watch"`
}

type Watch struct {
	File       string      `toml:"file"`
	Conditions []Condition `toml:"condition"`
}

type Condition struct {
	Matches      []string   `toml:"matches"`
	Threshold    int        `toml:"threshold"`
	WindowLength int        `toml:"window_length"`
	Comparison   Comparison `toml:"comparison"`
}

func ReadConfig(path string) (cfg *Config) {
	cfg = &Config{}
	if _, err := toml.DecodeFile(path, cfg); err != nil {
		log.Fatalf("Failed to decode config file: %s - %s", path, err)
	}
	return
}
