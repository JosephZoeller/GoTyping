package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/JosephZoeller/project-0/stringstats"
	"github.com/JosephZoeller/project-0/timer"
)

func preface() {
	fmt.Println("Welcome to my typing speed test. This program will count down from 3, and then it will measure how fast you can type words. " +
		"When you're ready, press the return key to begin and end.")
	read()
	timer.CountDown(3)
}

func runTest() {
	timer.ResetStopWatch()
	timer.BeginStopWatch()
	inputString, readError := read()
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
	str, err := reader.ReadString('\n')
	return str, err
}
