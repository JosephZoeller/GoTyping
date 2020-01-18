package termboxutil

import (
	"errors"
	"fmt"
	"github.com/nsf/termbox-go"
	"testing"
)

func TestWrite(t *testing.T) {
	str := "F"
	for i := 0; i < 100; i++ {
		str += "F"
		x, y := Write(0, 0, termbox.ColorDefault, termbox.ColorDefault, str)
		if x < 0 {
			t.Fatal(errors.New("x is less than 0"))
		} else if y < 0 {
			t.Fatal(errors.New("y is less than 0"))
		}
	}
}

func ExampleWrite() {
	str := "Example String"
	x, y := Write(0,0, termbox.ColorDefault, termbox.ColorDefault, str)
	fmt.Printf("x: %d, y: %d", x, y)
	//Output: x: 6, y: 1
}

func TestCountDown(t *testing.T) {
	ch := make(chan bool)
	defer close(ch)
	CountDown(0, 0, 5, "Counting down from %d...", ch)
	go CountDown(0, 0, 5, "Counting down from %d...", ch)
	ch <- true
}

func ExampleCountDown() {

	ch := make(chan bool)
	defer close(ch)
	CountDown(0, 0, 5, "Counting down from %d...", ch)
	// 5, 4, 3, 2, 1, 0
	go CountDown(0, 0, 5, "Counting down from %d...", ch)
	ch <- true
	// Returns immediately
}
