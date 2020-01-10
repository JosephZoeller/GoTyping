package stringstats

import (
	"fmt"
	"testing"
)

func TestCountWords(t *testing.T) {
	str := "Good morning, world! Goodbye, moon. TheSunIs  Very brihgt!"
	i := CountWords(str)
	if (i != 8) {
		t.Errorf("CountWords() miscounted.\nExpected count: %d,\nReturned count: %d", 8, i)
	}
}

func ExampleCountWords() {
	str := "Good morning, world! Goodbye, moon. TheSunIs  Very brihgt!"
	fmt.Println(CountWords(str))
	//Output: 8
}