package main

import (
	//"github.com/JosephZoeller/project-0/timer"

	"strconv"
	"time"

	"github.com/nsf/termbox-go"
)

const consoleWidth int = 60
const coldef termbox.Attribute = termbox.ColorDefault

func tbMessage(xCell int, yCell int, textWidth int, foreColor termbox.Attribute, backColor termbox.Attribute, message string) {
	i := xCell
	for _, char := range message {
		termbox.SetCell(i, yCell, char, foreColor, backColor)
		if char == '\n' || i > textWidth && char == ' ' {
			yCell++
			i = xCell
		} else {
			i++
		}
	}
	termbox.Flush()
}

func preface() {
	termbox.Clear(coldef, coldef)
	c := 3
	s := "Welcome to my typing speed test, " + *user +
		". This program will count down from " + strconv.Itoa(c) +
		", and then it will measure how fast you can type words." +
		"\n" + "When you're ready, press any key to begin..."
	tbMessage(0, 0, consoleWidth, termbox.ColorBlue, coldef, s)
	termbox.Flush()
	pressAnyKey()
}

func tbCountDown(n int) {
	termbox.Clear(coldef, coldef)
	for n > 0 {
		tbMessage(0, 0, consoleWidth, coldef, coldef, strconv.Itoa(n))
		time.Sleep(time.Second)
		n--
	}
	tbMessage(0, 0, consoleWidth, coldef, coldef, "Go!")
}

func pressAnyKey() rune {

	for {
		ev := termbox.PollEvent()
		switch ev.Type {
		case termbox.EventKey:
			return ev.Ch
		}
	}
}

func readWord() string {
	str := ""
	readLoop:
	for {
		ev := termbox.PollEvent()
		switch ev.Type {
		case termbox.EventKey:
			if ev.Key == termbox.KeyEsc {
				break readLoop
			}
			tbMessage(0,1,consoleWidth, coldef, coldef, strconv.QuoteRune(ev.Ch))
			str += strconv.QuoteRune(ev.Ch)

		}
	}
	return str
}