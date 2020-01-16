package termboxutil

import (
	"github.com/JosephZoeller/project-0/timer"

	"github.com/nsf/termbox-go"
)

var snt string = ""
var wordHistory []string
var crntwrd = ""
var keyevent = ""

func KeyContinue(reqEnter bool) rune {

	for {
		ev := termbox.PollEvent()
		if ev.Type == termbox.EventKey {
			if !reqEnter || ev.Key == termbox.KeyEnter {
				termbox.Clear(COLDEF, COLDEF)
				return ev.Ch
			}
		}
	}
}

func Readln(sdur int) []string {
	t := 0.00
	wordHistory = make([]string, 0)
mainLoop: // logic heavily inspired by editbox.go from the termbox-go _demos
	for {
		evt := termbox.PollEvent()
		t, _ = timer.CheckStopWatch()

		if t >= 0.1 {
			switch evt.Type {
			case termbox.EventKey:
				switch evt.Key {
				case termbox.KeyEsc: // TODO cause it to skip to end
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
		if t >= float64(sdur) {
			//newLine() // assumes the user is finished with the word they were working on. requires additional logic to disregard the word during discrepancy check
			break mainLoop
		}
	}
	return wordHistory
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
		wordHistory = append(wordHistory, crntwrd)
	}
	snt = ""
	crntwrd = ""
}

func space() {
	if len(snt) <= 0 {
		return
	} else if len(crntwrd) > 0 {
		wordHistory = append(wordHistory, crntwrd)
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
		if len(wordHistory) > 0 {
			crntwrd = wordHistory[len(wordHistory)-1]
		} else {
			crntwrd = ""
		}
		wordHistory = wordHistory[:len(wordHistory)-1]
	} else if len(crntwrd) > 0 {
		crntwrd = crntwrd[:len(crntwrd)-1]
	}
	snt = snt[:lensnt-1]
}
