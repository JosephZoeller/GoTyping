package timer

import (
	"testing"
)

func TestCountDown(t *testing.T) {
	n := 5
	CountDown(n)
}

func ExampleCountDown() {
	CountDown(3)
	// Output: 3
	// 2
	// 1
	// Go!
}
