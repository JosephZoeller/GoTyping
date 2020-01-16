package main

import (
	"strconv"

	tbutil "github.com/JosephZoeller/project-0/termboxutil"
	tb "github.com/nsf/termbox-go"
)

func showPreface(username string) {
	tb.Clear(tbutil.COLDEF, tbutil.COLDEF)
	cd := 3
	pre := "Welcome to my typing speed test, " + username +
		". This program will count down from " + strconv.Itoa(cd) +
		", and then it will measure how fast you can type words." +
		"\n" + "When you're ready, press any key to begin..."

	tbutil.Write(0, 0, tb.ColorBlue, tbutil.COLDEF, pre)
	tbutil.KeyContinue(false)

	tbutil.CountDown(0, 0, 3, "%s")
}
