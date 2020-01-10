package timer

import (
	"fmt"
	"testing"
	"time"
)

func TestStopWatch(t *testing.T) {
	BeginStopWatch()
	time.Sleep(time.Second)
	PauseStopWatch()
	n, r := CheckStopWatch()
	if n < 1 {
		t.Errorf("Trial 1: Stopwatch returned %f after 1 second of beginning", n)
	} else if r == true {
		t.Error("Trial 1: Stopwatch is running when it should be paused")
	}
	CheckStopWatch()

	e := PauseStopWatch()
	if e == nil {
		t.Error("Trial 2: Two calls to PauseStopWatch() did not detect an error")
	}
	BeginStopWatch()
	e = BeginStopWatch()
	if e == nil {
		t.Error("Trial 2: Two calls to BeginStopWatch() did not detect an error")
	}
	time.Sleep(time.Second)
	n, r = CheckStopWatch()
	if n < 2 {
		t.Errorf("Trial 3: Stopwatch returned %f after 2 seconds of beginning", n)
	} else if r == false {
		t.Error("Trial 3: Stopwatch is paused when it should be running")
	}
	n2, r := CheckStopWatch()
	if n2 < n {
		t.Error("Trial 4: Stopwatch returned less time than what has previously been recorded during the current stopwatch session")
	} else if r == false {
		t.Error("Trial 4: Stopwatch is paused when it should be running")
	}

	PauseStopWatch()
	ResetStopWatch()
	n, r = CheckStopWatch()
	if n != 0 {
		t.Errorf("Trial 5: stopwatch is not resetting")
	} else if r == true {
		t.Error("Trial 5: Stopwatch is still running when it should be paused")
	}

}

func ExampleBeginStopWatch() {

	BeginStopWatch()
	time.Sleep(time.Second)

	t, r := CheckStopWatch()
	fmt.Printf("Time: %.1f, Stopwatch is running: %t", t, r)
	PauseStopWatch()
	ResetStopWatch()
	//Output: Time: 1.0, Stopwatch is running: true
}

func ExampleCheckStopWatch() {
	BeginStopWatch()
	time.Sleep(time.Second)
	t, r := CheckStopWatch()
	fmt.Printf("Time: %.1f, Stopwatch is running: %t\n", t, r)
	PauseStopWatch()
	t, r = CheckStopWatch()
	fmt.Printf("Time: %.1f, Stopwatch is running: %t\n", t, r)
	ResetStopWatch()
	t, r = CheckStopWatch()
	fmt.Printf("Time: %.1f, Stopwatch is running: %t\n", t, r)
	//Output: Time: 1.0, Stopwatch is running: true
	// Time: 1.0, Stopwatch is running: false
	// Time: 0.0, Stopwatch is running: false
}

func ExamplePauseStopWatch() {
	BeginStopWatch()
	time.Sleep(time.Second)
	PauseStopWatch()
	time.Sleep(time.Second)

	t, r := CheckStopWatch()
	fmt.Printf("Time: %.1f, Stopwatch is running: %t\n", t, r)
	ResetStopWatch()
	//Output: Time: 1.0, Stopwatch is running: false
}

func ExampleResetStopWatch() {
	BeginStopWatch()
	time.Sleep(time.Second)
	PauseStopWatch()
	t, r := CheckStopWatch()
	fmt.Printf("Time: %.1f, Stopwatch is running: %t\n", t, r)
	ResetStopWatch()
	t, r = CheckStopWatch()
	fmt.Printf("Time: %.1f, Stopwatch is running: %t\n", t, r)
	//Output: Time: 1.0, Stopwatch is running: false
	// Time: 0.0, Stopwatch is running: false
}
