package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/JosephZoeller/project-0/timer"
	"github.com/nsf/termbox-go"
)

var user *string
var freestyle *bool
var duration *string

func init() {
	user = flag.String("user", strings.Title(os.Getenv("USER")), "Profile name: defaults to the Operating System's current username.")
	freestyle = flag.Bool("freestyle", true, "Freestyle testing: the user has no writing prompt and can type whatever they please.")
	duration = flag.String("duration", "1:00", "Test duration: The length of time that the typing test will last. Format as <Minutes>:<Seconds>")
	flag.Parse()
}

func runTest() {
	timer.ResetStopWatch()
	tbCountDown(3)
	timer.BeginStopWatch()
	c := readLoop()
	timer.PauseStopWatch()
	t, _ := timer.CheckStopWatch()
	tbprintStats(c, t) // get count
	pressAnyKey()
}

func tbprintStats(c int, t float64) {
	cfl := float64(c)
	termbox.Clear(coldef, coldef)
	tbMessage(0, 0, consoleWidth, termbox.ColorBlue, coldef, fmt.Sprintf("Words written: %d", c))
	tbMessage(0, 1, consoleWidth, termbox.ColorBlue, coldef, fmt.Sprintf("Seconds to complete: %f", t))
	tbMessage(0, 2, consoleWidth, termbox.ColorBlue, coldef, fmt.Sprintf("Words per second: %f", cfl/t))
	tbMessage(0, 3, consoleWidth, termbox.ColorBlue, coldef, fmt.Sprintf("Words per minute: %f", cfl/t*60))
}
