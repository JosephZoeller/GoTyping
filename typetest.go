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
	duration = flag.String("duration", "1:23", "Test duration: The length of time that the typing test will last. Format as <Minutes>:<Seconds>")
	flag.Parse()

}

func runTest() {
	timer.ResetStopWatch()
	tbCountDown(3)
	tbMessage(0, 0, 12, coldef, coldef, "Start typing!", false)
	timer.BeginStopWatch()
	// countdown as goroutine
	cd, _ := timer.ParseCountDown(*duration)
	tbMessage(30, 0, 80, termbox.ColorBlue, coldef, fmt.Sprintf("%d",cd), false)
	//
	cw, lw  := readLoop()
	timer.PauseStopWatch()
	t, _ := timer.CheckStopWatch()
	tbprintStats(cw, lw, t) // get count
	pressAnyKey()
}

func tbprintStats(c, l int, t float64) {
	cfl := float64(c)
	lfl := float64(l)
	termbox.Clear(coldef, coldef)
	width, _ := termbox.Size()

	tbMessage(0, 0, width, termbox.ColorBlue, coldef, fmt.Sprintf("Seconds to complete: %f", t), false)

	tbMessage(0, 2, width, termbox.ColorBlue, coldef, fmt.Sprintf("Words written: %d", c), false)
	tbMessage(0, 3, width, termbox.ColorBlue, coldef, fmt.Sprintf("Words per second: %f", cfl/t), false)
	tbMessage(0, 4, width, termbox.ColorBlue, coldef, fmt.Sprintf("Words per minute: %f", cfl/t*60), false)

	tbMessage(30, 2, width, termbox.ColorBlue, coldef, fmt.Sprintf("Characters written: %d", l), false)
	tbMessage(30, 3, width, termbox.ColorBlue, coldef, fmt.Sprintf("Characters per second: %f", lfl/t), false)
	tbMessage(30, 4, width, termbox.ColorBlue, coldef, fmt.Sprintf("Characters per minute: %f", lfl/t*60), false)
}
