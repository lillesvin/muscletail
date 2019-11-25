package main

import (
	log "github.com/sirupsen/logrus"
	"os/exec"
	"strings"
)

func RunAction(cmd string) bool {
	sc := strings.Split(cmd, ` `)

	var c *exec.Cmd
	if len(sc) == 1 {
		c = exec.Command(sc[0])
	} else {
		c = exec.Command(sc[0], sc[1:]...)
	}

	err := c.Run()
	if err != nil {
		log.WithFields(log.Fields{
			"cmd": cmd,
		}).Warn("Failed to run action")
		return false
	}
	return true
}
