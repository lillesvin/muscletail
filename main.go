package main

import (
	log "github.com/sirupsen/logrus"
)

func main() {
	//cfg := ReadConfig("config/config.default.toml")

	w := &Window{Length: 10}
	w.Append(1)
	w.Append(2)
	w.Append(3)
	w.Append(4)
	w.Append(5)
	w.Append(6)
	w.Append(7)
	w.Append(8)
	w.Append(9)
	w.Append(10)
	w.Append(11)
	w.Append(12)
	log.Infof("%+v", w)
}
