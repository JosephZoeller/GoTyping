package timer

import (
	"errors"
	"time"
)

var startTime time.Time
var totalTime time.Duration = 0
var isTiming bool = false

//BeginStopWatch starts the stopwatch.
func BeginStopWatch() error {
	if isTiming {
		return errors.New("The stopwatch is already on")
	}
	isTiming = true
	startTime = time.Now()
	return nil
}

//PauseStopWatch temporarily stops the stopwatch.
func PauseStopWatch() error {
	if !isTiming {
		return errors.New("The stopwatch is not on")
	}
	isTiming = false
	totalTime = totalTime + time.Since(startTime)
	return nil
}

//CheckStopWatch returns the elapsed time in seconds that the stopwatch has been running. It can return the elapsed time even if the stopwatch is paused.
func CheckStopWatch() float64 {
	if isTiming {
		return time.Since(startTime).Seconds() + totalTime.Seconds()
	}
	return totalTime.Seconds()
}

// ResetStopWatch resets the stopwatch, restarting the elapsed time.
func ResetStopWatch() {
	startTime = time.Now()
	totalTime = 0
}
