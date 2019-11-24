package main

func main() {
	cfg := ReadConfig("config/config.default.toml")

	for _, w := range cfg.Watches {
		go w.Monitor()
	}

	// Wait forever
	select {}
}
