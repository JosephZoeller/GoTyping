package termboxutil

import (
	"github.com/JosephZoeller/project-0/timer"
	"log"

	"github.com/nsf/termbox-go"
)

var snt string = ""
var wordHistory []string
var crntwrd = ""
var keyevent = ""
var totalWords int

func KeyContinue(reqEnter bool) rune {
	log.Printf("[termboxutil]: Awaiting keypress (Enter key required: %t)...", reqEnter)
	for {
		ev := termbox.PollEvent()
		if ev.Type == termbox.EventKey {
			if !reqEnter || ev.Key == termbox.KeyEnter {
				termbox.Clear(COLDEF, COLDEF)
				log.Printf("[termboxutil]: Keypress accepted")
				return ev.Ch
			}
		}
	}
}

func Readln(sdur int, verb bool) []string {
	log.Printf("[termboxutil]: Reading user input")
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
			if verb {
				drawRTStats()
			}
		}
		if t >= float64(sdur) {
			//newLine() // assumes the user is finished with the word they were working on. would require additional logic to disregard the word during discrepancy check
			totalWords = 0
			break mainLoop
		}
	}
	log.Printf("[termboxutil]: Read loop ended:\n\ttime: %.2f\n\twordHistory: %s\n\ttotalWords: %d", t, wordHistory, totalWords)
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
		totalWords++
	}
	snt = ""
	crntwrd = ""
}

func space() {
	if len(snt) <= 0 {
		return
	} else if len(crntwrd) > 0 {
		wordHistory = append(wordHistory, crntwrd)
		totalWords++
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
		totalWords--
		wordHistory = wordHistory[:len(wordHistory)-1]
	} else if len(crntwrd) > 0 {
		crntwrd = crntwrd[:len(crntwrd)-1]
	}
	snt = snt[:lensnt-1]
}
