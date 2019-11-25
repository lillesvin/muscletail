package main

import (
	"os"
	"runtime"
	"strings"
	"time"

	"github.com/hpcloud/tail"
	log "github.com/sirupsen/logrus"
)

// Watch defines one or more things to look for in a given file
// along with window properties and trigger threshold
// TODO: Support defining an action here
type Watch struct {
	ID           string     `toml:"id"`
	File         string     `toml:"file"`
	Matches      []string   `toml:"matches"`
	Threshold    int        `toml:"threshold"`
	WindowLength int        `toml:"window_length"`
	Comparison   Comparison `toml:"comparison"`
	Action       string     `toml:"action"`
}

// Comparison type for use in constants below
type Comparison string

const (
	// GreaterThan means comparison by >
	GreaterThan Comparison = "GreaterThan"
	// LessThan compares by <
	LessThan Comparison = "LessThan"
)

// Monitor sets up a tail on a file and continuously
// checks against the defined conditions if an action
// should be triggered
func (w *Watch) Monitor() {
	// Default to a window length of 10
	if w.WindowLength < 1 {
		w.WindowLength = 10
	}

	// Open a new window
	win := &Window{Length: w.WindowLength}
	mark := time.Now()

	// Get some tail
	t := w.setupTail()

	// Tell all about it
	w.announceNew()

	// This is basically the main loop
	// We get a line and possibly react to it
	for l := range t.Lines {
		if w.stringSearch(l.Text) {
			win.Append(time.Since(mark).Seconds())
			mark = time.Now()

			// Don't trigger unless window is full and conditions are met
			if win.Full() && w.conditionsMet(win.WindowMean) {
				w.triggerAction(win)
			}
		}
	}
}

func (w *Watch) setupTail() *tail.Tail {
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

	// Use Polling on Windows
	if runtime.GOOS == "windows" {
		t.Config.Poll = true
	}

	if err != nil {
		log.Fatalf("Error tailing '%s': %s", w.File, err)
	}

	return t
}

func (w *Watch) announceNew() {
	log.WithFields(log.Fields{
		"ID":           w.ID,
		"File":         w.File,
		"Matches":      w.Matches,
		"Threshold":    w.Threshold,
		"Comparison":   w.Comparison,
		"WindowLength": w.WindowLength,
	}).Info("New watch")
}

func (w *Watch) conditionsMet(val float64) bool {
	switch w.Comparison {
	case GreaterThan:
		return val > float64(w.Threshold)
	case LessThan:
		return val < float64(w.Threshold)
	}

	// Default to LessThan
	return val < float64(w.Threshold)
}

func (w *Watch) stringSearch(haystack string) bool {
	for _, needle := range w.Matches {
		if strings.Contains(haystack, needle) {
			return true
		}
	}
	return false
}

func (w *Watch) triggerAction(win *Window) {
	log.WithFields(log.Fields{
		"ID":           w.ID,
		"File":         w.File,
		"WindowMean":   win.WindowMean,
		"WindowLength": w.WindowLength,
		"Threshold":    w.Threshold,
	}).Warn("Achtung, baby!")

	RunAction(w.Action)
}
