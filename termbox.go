package main

import (
	"github.com/JosephZoeller/project-0/timer"

	//"fmt"
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

var cursor int

func readSentence() string {
	snt := ""
readSentenceLoop:
	for {
		s, ev := readWord()
		switch ev.Type {
		case termbox.EventKey:
			switch ev.Key {
			case termbox.KeyEnter, termbox.KeyEsc:
				snt += s
				break readSentenceLoop
			case termbox.KeySpace:
				snt += s + " "
			}
		}
	}
	tbMessage(0, 0, consoleWidth, coldef, coldef, snt)
	return snt
}

func readWord() (string, termbox.Event) {
	r := make([]rune, 0)
	var swt float64
	var ev termbox.Event
readWordLoop:
	for {
		ev = termbox.PollEvent()
		switch ev.Type {
		case termbox.EventKey:
			switch ev.Key {
			case termbox.KeySpace, termbox.KeyEsc, termbox.KeyEnter:
				break readWordLoop
			default:
				tbMessage(0, 1, consoleWidth, coldef, coldef, strconv.QuoteRune(ev.Ch))
				r = append(r, ev.Ch)
			}
		}
	}
	// If the user has input entire words within a tenth of a second, it's reasonable to assume that they had preemptively typed something during the countdown.
	// this a kludge but I haven't figured out a way of denying or ignoring user input during the countdown, so in the interest of time, this is the best alternative.
	swt, _ = timer.CheckStopWatch()
	if swt < 0.1 {
		r = make([]rune, 0)
		goto readWordLoop
	}
	return string(r), ev
}
