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
	duration = flag.String("duration", "0:11", "Test duration: The length of time that the typing test will last. Format as <Minutes>:<Seconds>")
	flag.Parse()

}

func runTest() {
	tbCountDown(3, 0, 0, "%s")
	cd, _ := timer.ParseCountDown(*duration)
	timer.BeginStopWatch()
	tbMessage(0, 0, coldef, coldef, "Start typing!")
	go tbCountDown(cd, 50, 0, "Time Remaining: %s Seconds...")
	cw, lw := readLoop(cd)
	timer.PauseStopWatch()
	t, _ := timer.CheckStopWatch()
	tbprintStats(cw, lw, t) // get count
	keyContinue(true)
}

func tbprintStats(c, l int, t float64) {
	cfl := float64(c)
	lfl := float64(l)
	termbox.Clear(coldef, coldef)
	termbox.HideCursor()
	tbMessage(0, 0, termbox.ColorBlue, coldef, fmt.Sprintf("Seconds to complete: %f", t))

	tbMessage(0, 2, termbox.ColorBlue, coldef, fmt.Sprintf("Words written: %d", c))
	tbMessage(0, 3, termbox.ColorBlue, coldef, fmt.Sprintf("Words per second: %f", cfl/t))
	tbMessage(0, 4, termbox.ColorBlue, coldef, fmt.Sprintf("Words per minute: %f", cfl/t*60))

	tbMessage(30, 2, termbox.ColorBlue, coldef, fmt.Sprintf("Characters written: %d", l))
	tbMessage(30, 3, termbox.ColorBlue, coldef, fmt.Sprintf("Characters per second: %f", lfl/t))
	tbMessage(30, 4, termbox.ColorBlue, coldef, fmt.Sprintf("Characters per minute: %f", lfl/t*60))

	tbMessage(0, 6, coldef, coldef, "Press the enter key to end the program...")
}
