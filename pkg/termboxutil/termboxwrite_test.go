// Package termboxutil project-0 Typing Test for 200106-uta-go
package termboxutil

import (
	"fmt"
	"github.com/nsf/termbox-go"
	"testing"
)

func BenchmarkWrite(b *testing.B) {
	str := "F "
	for i := 0; i < b.N; i++ {
		str += "F "
		Write(0, 0, termbox.ColorDefault, termbox.ColorDefault, str)
	}
}

func TestWrite(t *testing.T) {
	str := "This is a test."
	x, y := Write(0, 0, termbox.ColorDefault, termbox.ColorDefault, str)
	if x != 5 {
		t.Errorf("X value is off by %d", 5 - x)
	} else if y != 3 {
		t.Errorf("Y value is off by %d", 3 - y)
	}
}

func ExampleWrite() {
	str := "1234 123 12345"
	width, _ := termbox.Size()                                           // In the absence of a terminal, returns 0, 0
	x, y := Write(0, 0, termbox.ColorDefault, termbox.ColorDefault, str) // If Terminal is 0, prints 1 word per line
	fmt.Printf("Terminal Width: %d. x: %d, y: %d", width, x, y)
	//Output: Terminal Width: 0. x: 5, y: 2
}

func TestCountDown(t *testing.T) {
	ch := make(chan bool, 2)
	defer close(ch)
	e := CountDown(0, 0, -1, "Counting down from %d", ch)
	if e == nil {
		t.Errorf("CountDown accepted negative input")
	}
	e = CountDown(0, 0, 5, "Counting down from %d...", ch)
	if e != nil {
		t.Errorf("CountDown rejected a valid input")
	}
	ch <- true
	e = CountDown(0, 0, 5, "Counting down from %d...", ch)
	if e == nil {
		t.Errorf("CountDown was not aborted")
	}
}

func ExampleCountDown() {

	ch := make(chan bool)
	defer close(ch)
	CountDown(0, 0, 5, "Counting down from %d...", ch)
	// Output: 5, 4, 3, 2, 1, 0
	go CountDown(0, 0, 5, "Counting down from %d...", ch)
	ch <- true
	// Immediately Aborts CountDown
}
