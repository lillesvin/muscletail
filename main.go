package main

import (
	"github.com/hpcloud/tail"
	log "github.com/sirupsen/logrus"
	"os"
	"strings"
	"time"
)

func main() {
	cfg := ReadConfig("config/config.default.toml")

	for _, w := range cfg.Watches {
		go Monitor(w)
	}

	// Wait forever
	select {}
}

func Monitor(w Watch) {
	log.Infof("New watch: %+v", w)

	win := &Window{Length: w.WindowLength}
	mark := time.Now()

	t, err := tail.TailFile(w.File, tail.Config{
		Follow:    true,
		ReOpen:    true,
		MustExist: true,
		Logger:    tail.DiscardingLogger,
		Location: &tail.SeekInfo{
			Offset: 0,
			Whence: os.SEEK_END,
		},
	})
	if err != nil {
		log.Fatalf("Error tailing '%s': %s", w.File, err)
	}

	log.Infof("Watching '%s'", w.File)

	for l := range t.Lines {
		if StringSearch(l.Text, w.Matches) {
			win.Append(time.Since(mark).Seconds())
			mark = time.Now()
			if win.Full() && MeetsConditions(win, w) {
				Trigger(win, w)
			}
		}
	}
}

func MeetsConditions(win *Window, w Watch) bool {
	switch w.Comparison {
	case GreaterThan:
		return win.WindowMean > float64(w.Threshold)
	case LessThan:
		return win.WindowMean < float64(w.Threshold)
	}

	// Default to LessThan
	return win.WindowMean < float64(w.Threshold)
}

func StringSearch(haystack string, needle []string) bool {
	for _, n := range needle {
		if strings.Contains(haystack, n) {
			return true
		}
	}
	return false
}

func Trigger(win *Window, w Watch) {
	log.Printf("%s [%s] (Value: %.2f, threshold: %d)", w.ID, w.File, win.WindowMean, w.Threshold)
}
