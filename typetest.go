package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/JosephZoeller/project-0/stringstats"
	"github.com/JosephZoeller/project-0/timer"
)

var user *string
var freestyle *bool
var duration *string

func init() { //WIP
	user = flag.String("user", strings.Title(os.Getenv("USER")), "Profile name: defaults to the Operating System's current username.")
	freestyle = flag.Bool("freestyle", true, "Freestyle testing: the user has no writing prompt and can type whatever they please.")
	duration = flag.String("duration", "1:00", "Test duration: The length of time that the typing test will last. Format as <Minutes>:<Seconds>")
	flag.Parse()

	fmt.Println(*user, *freestyle, *duration)
}

func preface() { // Intro to the test, prior to it starting.
	fmt.Printf("Welcome to my typing speed test, %s. This program will count down from 3, and then it will measure how fast you can type words. "+
		"When you're ready, press the return key to begin and end.\n", *user)
	timer.CountDown(3)

}

func runTest() {
	timer.ResetStopWatch()
	timer.BeginStopWatch()
	inputString, readError := read() // BUG: If the user starts to write to stdin prior to the countdown ending, it's counted toward the test score/statistics, resulting in unrealistically fast typing speeds.
	timer.PauseStopWatch()
	elapsedTime, _ := timer.CheckStopWatch()

	if readError != nil {
		fmt.Println(readError)
		os.Exit(-1)
	}

	printStats(elapsedTime, stringstats.CountWords(inputString))
}

func printStats(seconds float64, wordCount int) {
	wordCountFl := float64(wordCount)
	fmt.Println("Seconds to complete:", seconds)
	fmt.Println("Words per second:", wordCountFl/seconds)
	fmt.Println("Words per minute:", wordCountFl/seconds*60)
}

func read() (string, error) {
	reader := bufio.NewReader(os.Stdin)
	reader.Reset(os.Stdin) // not working the way I was hoping it would
	str, err := reader.ReadString('\n')
	return str, err
}
