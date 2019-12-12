package main

import (
	"flag"
	log "github.com/sirupsen/logrus"
	"os"
)

var configFile string
var defaultConfig []string

func init() {
	defaultConfig = []string{
		"./.muscletail.toml",
		"./muscletail.toml",
		"/etc/muscletail.toml",
	}

	flag.StringVar(&configFile, "config", "", "Use this config")
	flag.StringVar(&configFile, "c", "", "Shorthand for -config")
	flag.Parse()
}

func main() {
	if configFile == "" {
		for _, f := range defaultConfig {
			info, err := os.Stat(f)
			if !os.IsNotExist(err) && !info.IsDir() {
				configFile = f
				break
			}
		}
	}

	if configFile == "" {
		log.Fatal("No config found")
	}
	log.WithFields(log.Fields{
		"Config": configFile,
	}).Info("Found config file")

	cfg := ReadConfig(configFile)

	for _, w := range cfg.Watches {
		go w.Monitor()
	}

	// Wait forever
	select {}
}
