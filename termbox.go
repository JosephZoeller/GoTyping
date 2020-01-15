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

func tbMessage(xCell int, yCell int, foreColor termbox.Attribute, backColor termbox.Attribute, message string) (int, int) {
	i := xCell
	textWidth, _ := termbox.Size()
	for _, char := range message {
		termbox.SetCell(i, yCell, char, foreColor, backColor)
		if i > textWidth && char == ' ' {
			yCell++
			i = xCell
		} else {
			i++
		}
	}
	for j := i; j < textWidth; j++ {
		termbox.SetCell(j, yCell, ' ', foreColor, backColor)
	}
	termbox.Flush()
	return i, yCell
}

func preface() {
	termbox.Clear(coldef, coldef)
	c := 3
	s := "Welcome to my typing speed test, " + *user +
		". This program will count down from " + strconv.Itoa(c) +
		", and then it will measure how fast you can type words." +
		"\n" + "When you're ready, press any key to begin..."
	tbMessage(0, 0, termbox.ColorBlue, coldef, s)
	keyContinue(false)
}

func tbCountDown(n, x, y int, frmt string) {
	for n > 0 {
		tbMessage(x, y, coldef, coldef, fmt.Sprintf(frmt, strconv.Itoa(n)))
		time.Sleep(time.Second)
		n--
	}
	tbMessage(x, y, coldef, coldef, fmt.Sprintf(frmt, "0"))
}

func keyContinue(reqEnter bool) rune {

	for {
		ev := termbox.PollEvent()
		if ev.Type == termbox.EventKey {
			if !reqEnter || ev.Key == termbox.KeyEnter {
				termbox.Clear(coldef, coldef)
				return ev.Ch
			}
		}
	}
}

var snt string = ""
var wrds = make([]string, 0)
var crntwrd = ""
var keyevent = ""

func readLoop(sdur int) (int, int) {
	t := 0.00
mainLoop: // logic heavily inspired by termbox-go demo, editbox.go
	for {
		evt := termbox.PollEvent()
		t, _ = timer.CheckStopWatch()
		if t >= float64(sdur) {
			newLine()
			break mainLoop
		}
		if t >= 0.1 {
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
	cl := len(wrds) - 1 //# spaces between words
	for _, wrd := range wrds {
		cl += len(wrd)
	}
	return len(wrds), cl
}

func readLoopSentence(sdur int) (int, int, []string) {
	t := 0.00
mainLoop: // logic heavily inspired by termbox-go demo, editbox.go
	for {
		evt := termbox.PollEvent()
		t, _ = timer.CheckStopWatch()
		if t >= float64(sdur) {
			newLine()
			break mainLoop
		}
		if t >= 0.1 {
			switch evt.Type {
			case termbox.EventKey:
				switch evt.Key {
				case termbox.KeyEsc:
					break mainLoop
				case termbox.KeyEnter:
					newLine()
					redraw()
					keyevent = "[Enter]"
					break mainLoop
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
	cl := len(wrds) - 1 //# spaces between words
	for _, wrd := range wrds {
		cl += len(wrd)
	}
	return len(wrds), cl, wrds
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
	} else if len(crntwrd) > 0 {
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
	sntX, sntY := tbMessage(0, 1, coldef, coldef, snt)
	termbox.SetCursor(sntX, sntY)

	tbMessage(0, 3, coldef, coldef, ("Event: " + keyevent))
	tbMessage(0, 4, coldef, coldef, fmt.Sprintf("Word Bank: %s", wrds))
	tbMessage(0, 5, coldef, coldef, "Current word: "+crntwrd)
}
