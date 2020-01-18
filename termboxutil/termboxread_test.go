package termboxutil

import (
)

func ExampleReadln() {
	Readln(30, true)
	// listens and displays user input for 30 seconds as well as verbose details. Enter or Escape key will end the listening.
	// returns the words (characters delimited by a ' ') input by the user as a slice of strings.
	Readln(10, false)
	// listens and displays for 10 seconds, with no verbose details
}

func ExampleKeyContinue() {
	KeyContinue(false)
	// awaits user input.
	KeyContinue(true)
	// awaits enter keypress
}