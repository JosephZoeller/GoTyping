// Package timer project-0 Typing Test for 200106-uta-go
package timer

import (
	"errors"
	"log"
	"time"
)

var startTime time.Time
var totalTime time.Duration = 0
var isTiming bool = false

//BeginStopWatch starts the stopwatch.
func BeginStopWatch() error {
	log.Println("[stopwatch]: Timer started")
	if isTiming {
		log.Println("[stopwatch]: Error: The stopwatch is already on.")
		return errors.New("The stopwatch is already on")
	}
	isTiming = true
	startTime = time.Now()
	return nil
}

//PauseStopWatch temporarily stops the stopwatch.
func PauseStopWatch() error {
	if !isTiming {
		log.Println("[stopwatch]: Error: The stopwatch is not on.")
		return errors.New("The stopwatch is not on")
	}
	isTiming = false
	totalTime = totalTime + time.Since(startTime)
	log.Println("[stopwatch]: Timer paused: Elapsed " + totalTime.String())
	return nil
}

//CheckStopWatch returns the elapsed time in seconds that the stopwatch has been running. It can return the elapsed time even if the stopwatch is paused.
func CheckStopWatch() (float64, bool) {
	if isTiming {
		//log.Println("[stopwatch]: Timer checked: " + startTime.String()) // too frequent to be useful
		return time.Since(startTime).Seconds() + totalTime.Seconds(), isTiming
	}
	log.Println("[stopwatch]: Timer checked: Elapsed " + totalTime.String())
	return totalTime.Seconds(), isTiming
}

// ResetStopWatch resets the stopwatch, restarting the elapsed time.
func ResetStopWatch() {
	startTime = time.Now()
	totalTime = 0
	log.Println("[stopwatch]: Timer Reset")
}

// PrimeStopWatch stops and then resets the stopwatch.
func PrimeStopWatch() {
	_, running := CheckStopWatch()
	if running {
		PauseStopWatch()
	}
	ResetStopWatch()
}
