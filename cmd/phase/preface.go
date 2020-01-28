// Package phase project-0 Typing Test for 200106-uta-go
package phase

import (
	"fmt"

	tbutil "github.com/JosephZoeller/GoTyping/pkg/termboxutil"
	tb "github.com/nsf/termbox-go"
)

// ShowPreface displays an introduction message to the user before moving onto the typing test. Accepts a string for a username
// as well as the freestyle boolean, which will determine what briefing will be displayed to the user.
func ShowPreface(username string, duration int, free bool) {
	tb.Clear(tbutil.COLDEF, tbutil.COLDEF)
	cd := 3
	var freestr string
	var mode string
	if free {
		mode = "Freestyle"
		freestr = "For this mode, you aren't given a writing prompt. You can write whatever you please before the timer runs out." +
			" During the test, press the Enter key to commit your sentence and begin a new line. You can also press the escape key to end the test."
	} else {
		mode = "Prompt-Matching"
		freestr = "For this mode, you are given a writing prompt to copy. You should try to match the prompt exactly." +
			" During the test, press the Enter key to commit your sentence and recieve a new writing prompt. You can also press the escape key to end the test."
	}


	params := fmt.Sprintf("[Test Mode: %s, Duration: %ds]", mode, duration)
	pre := fmt.Sprintf("Welcome to my typing speed test, %s. This program will count down from %d, and then it'll measure how fast you can type words. %s"+
		" When you're ready, press any key to begin...", username, cd, freestr)

	_, y := tbutil.Write(0, 0, tbutil.COLDEF, tbutil.COLDEF, params)	
	tbutil.Write(0, y + 2, tbutil.COLDEF, tbutil.COLDEF, pre)

	tbutil.KeyContinue(false)

	tbutil.CountDown(0, 0, 3, "%s", nil)
}
