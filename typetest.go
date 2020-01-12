package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	//"github.com/JosephZoeller/project-0/stringstats"
	"github.com/JosephZoeller/project-0/timer"
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
	// do test
	readWord()
	timer.PauseStopWatch()
}

func printStats() {
	fmt.Println("Seconds to complete:")
	fmt.Println("Words per second:")
	fmt.Println("Words per minute:")
}
