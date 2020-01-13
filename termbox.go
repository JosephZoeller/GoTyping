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
	for i < textWidth {
		termbox.SetCell(i, yCell, ' ', foreColor, backColor)
		i++
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

var snt string = ""
var wrds = make([]string, 0)
var crntwrd = ""
var keyevent = ""

func readLoop() int {
	t := 0.00
mainLoop:
	for {
		evt := termbox.PollEvent()
		t, _ = timer.CheckStopWatch() //kludge: currently the only thing I can think of to keep users from preempting words into console during the countdown
		if t >= 0.1 {				  // TODO find proper solution
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
	return len(wrds)
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
	snt = ""
	crntwrd = ""
}

func space() {
	if len(snt) <= 0 {
		return
	}
	wrds = append(wrds, crntwrd)
	snt += " "
	crntwrd = ""
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
	tbMessage(0, 1, consoleWidth, coldef, coldef, snt)

	tbMessage(0, 3, consoleWidth, coldef, coldef, ("Event: " + keyevent))
	tbMessage(0, 4, consoleWidth, coldef, coldef, fmt.Sprintf("Word Bank: %s", wrds))
	tbMessage(0, 5, consoleWidth, coldef, coldef, "Current word: "+crntwrd)
}
