package main

import (
	"fmt"

	tbutil "github.com/JosephZoeller/project-0/termboxutil"
	tb "github.com/nsf/termbox-go"
)

func showPreface(username string, free bool) {
	tb.Clear(tbutil.COLDEF, tbutil.COLDEF)
	cd := 3
	var freestr string
	if (free) {
		freestr = "For this mode, you aren't given a writing prompt. You can write whatever you'd like before the timer runs out." +
		" During the test, press the Enter key to commit your sentence and begin a new line. You can also press the escape key to end the test."
	} else {
		freestr = "For this mode, you are given a writing prompt to copy. You should try to match the prompt exactly." +
		" During the test, press the Enter key to commit your sentence and recieve a new writing prompt. You can also press the escape key to end the test."
	}
	pre := fmt.Sprintf("Welcome to my typing speed test, %s. This program will count down from %d, and then it will measure how fast you can type words. %s" +
						" When you're ready, press any key to begin...", username, cd, freestr)
	tbutil.Write(0, 0, tb.ColorBlue, tbutil.COLDEF, pre)
	tbutil.KeyContinue(false)

	tbutil.CountDown(0, 0, 3, "%s", nil)
}
