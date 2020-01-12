package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	//"github.com/JosephZoeller/project-0/stringstats"
	"github.com/JosephZoeller/project-0/timer"
	//"github.com/nsf/termbox-go"
)

var user *string
var freestyle *bool
var duration *string

func init() {
	user = flag.String("user", strings.Title(os.Getenv("USER")), "Profile name: defaults to the Operating System's current username.")
	freestyle = flag.Bool("freestyle", true, "Freestyle testing: the user has no writing prompt and can type whatever they please.")
	duration = flag.String("duration", "1:00", "Test duration: The length of time that the typing test will last. Format as <Minutes>:<Seconds>")
	flag.Parse()

	fmt.Println(*user, *freestyle, *duration)
}

func preface() {
	c := 3
	fmt.Printf("Welcome to my typing speed test, %s. This program will count down from %d, and then it will measure how fast you can type words. "+
		"When you're ready, press the return key to begin and end.\n", *user, c)
	// read any key
	timer.CountDown(3)

}

func runTest() {
	timer.ResetStopWatch()
	timer.BeginStopWatch()
	// do test
	timer.PauseStopWatch()
}

func printStats() {
	fmt.Println("Seconds to complete:")
	fmt.Println("Words per second:")
	fmt.Println("Words per minute:")
}
