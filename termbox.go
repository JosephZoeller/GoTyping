package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/JosephZoeller/project-0/timer"

	"github.com/nsf/termbox-go"
)

const consoleWidth int = 60
const coldef termbox.Attribute = termbox.ColorDefault

func tbMessage(xCell int, yCell int, textWidth int, foreColor termbox.Attribute, backColor termbox.Attribute, message string, showCursor bool) {
	i := xCell
	for _, char := range message {
		termbox.SetCell(i, yCell, char, foreColor, backColor)
		if showCursor {
			termbox.SetCursor(i, yCell)
		} else {
			termbox.HideCursor() // will require cursor-enabled message to come last when drawing a number of messages 
		}
		if char == '\n' || i > textWidth && char == ' ' {
			yCell++
			i = xCell
		} else {
			i++
		}
	}
	for i < textWidth {
		termbox.SetCell(i, yCell, ' ', foreColor, backColor)
		i++
	}
	termbox.Flush()
}

func preface() {
	termbox.Clear(coldef, coldef)
	width, _ := termbox.Size()
	c := 3
	s := "Welcome to my typing speed test, " + *user +
		". This program will count down from " + strconv.Itoa(c) +
		", and then it will measure how fast you can type words." +
		"\n" + "When you're ready, press any key to begin..."
	tbMessage(0, 0, width, termbox.ColorBlue, coldef, s, true)
	termbox.Flush()
	pressAnyKey()
}

func tbCountDown(n int) {
	termbox.Clear(coldef, coldef)
	for n > 0 {
		tbMessage(0, 0, 1, coldef, coldef, strconv.Itoa(n), false)
		time.Sleep(time.Second)
		n--
	}
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

var snt string = ""
var wrds = make([]string, 0)
var crntwrd = ""
var keyevent = ""

func readLoop() (int, int) {
	t := 0.00
mainLoop: // logic heavily inspired by termbox-go demo, editbox.go
	for {
		evt := termbox.PollEvent()
		t, _ = timer.CheckStopWatch() //kludge: currently the only thing I can think of to keep users from preempting words into console during the countdown
		if t >= 0.1 {                 // TODO find proper solution
			switch evt.Type {
			case termbox.EventKey:
				switch evt.Key {
				case termbox.KeyEsc:
					break mainLoop
				case termbox.KeyEnter:
					newLine()
					keyevent = "[Enter]"
				case termbox.KeySpace:
					space()
					keyevent = "[Space]"
				case termbox.KeyBackspace2, termbox.KeyBackspace:
					backspace()
					keyevent = "[Backspace]"
				default:
					addRune(evt.Ch)
					keyevent = "[AddRune]"
				}
			case termbox.EventError:
				// expand
				break mainLoop
			}
			redraw()
		}
	}
	return len(wrds), len(snt)
}

func addRune(r rune) {
	s := string(r)
	crntwrd += s
	snt += s
}

func newLine() {
	if len(snt) <= 0 {
		return
	}
	if crntwrd != "" {
		wrds = append(wrds, crntwrd)
	}
	//snt = "" not sure if I need/want this
	crntwrd = ""
}

func space() {
	if len(snt) <= 0 {
		return
	} else if (len(crntwrd) > 0) {
		wrds = append(wrds, crntwrd)
		snt += " "
		crntwrd = ""
	}
}

func backspace() {
	lensnt := len(snt)
	if lensnt <= 0 {
		return
	}
	if snt[lensnt-1] == ' ' {
		if len(wrds) > 0 {
			crntwrd = wrds[len(wrds)-1]
		} else {
			crntwrd = ""
		}
		wrds = wrds[:len(wrds)-1]
	} else if len(crntwrd) > 0 {
		crntwrd = crntwrd[:len(crntwrd)-1]
	}
	snt = snt[:lensnt-1]
}

func redraw() {
	sntY := 1
	width, _  := termbox.Size()

	tbMessage(0, 3, width, coldef, coldef, ("Event: " + keyevent), false)
	tbMessage(0, 4, width, coldef, coldef, fmt.Sprintf("Word Bank: %s", wrds), false)
	tbMessage(0, 5, width, coldef, coldef, "Current word: "+crntwrd, false)

	tbMessage(0, sntY, width, coldef, coldef, snt, true)
}
