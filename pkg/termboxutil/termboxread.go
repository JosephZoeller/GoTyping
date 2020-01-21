// Package termboxutil project-0 Typing Test for 200106-uta-go
package termboxutil

import (
	"errors"
	"github.com/JosephZoeller/project-0/pkg/timer"
	"log"

	"github.com/nsf/termbox-go"
)
// wordHistory: Saves the user's previous words which are added and removed as the user types. The variable which is returned for analysis.
var wordHistory []string
// crntwrd: Current word. Pushes onto wordHistory when the space key is pressed. wordHistory pops off it's last word into crntwrd when the backspace removes a space.
var crntwrd = ""

// KeyContinue Pauses the process to await a key press from the users.
// The function accepts a bool, which determines if the Enter key needs to be the key press.
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

// Readln Reads user's keypresses in real-time. If the stopwatch is not running, it will return an error.
// Accepts a positive integer, which determines the duration that the function will listen.
// Also accepts a bool, which will display verbose analytics about the user input.
// Returns the user's input as a slice of strings. Each string represents a word (letters delimited by ' ') from their input.
// The function can be ended naturally prior to the time limit by pressing the Enter key.
// Pressing the Escape key will also end the function, but with an error message.
func Readln(sdur int, verb bool) ([]string, error) {
	var err error
	var t float64
	_, on := timer.CheckStopWatch()
	wordHistory = make([]string, 0)
	keyevent := ""

	if !on {
		log.Printf("[termboxutil]: Error: The stopwatch is not running and Readln cannot measure the duration.")
		err = errors.New("The stopwatch is not running and Readln cannot measure the duration.")
		return wordHistory, err
	} else if sdur <= 0 {
		log.Printf("[termboxutil]: Error: The Readln duration argument is less than or equal to 0")
		err = errors.New("The Readln duration argument is less than or equal to 0.")
		return wordHistory, err
	}
	redraw(verb, keyevent)
	log.Printf("[termboxutil]: Reading user input")
readLoop: // logic heavily inspired by editbox.go from the termbox-go _demos
	for {
		evt := termbox.PollEvent()
		t, _ = timer.CheckStopWatch()

		if t >= 0.1 {
			switch evt.Type {
			case termbox.EventKey:
				
				switch evt.Key {
				case termbox.KeyEsc:
					log.Printf("[termboxutil]: Exit: Encountered termbox.KeyEsc during Readln loop.")
					err = errors.New("termbox.KeyEsc Event during Readln loop.")
					break readLoop
				case termbox.KeyEnter:
					keyevent = "[Enter]"
					space()
					redraw(verb, keyevent)
					break readLoop
				case termbox.KeySpace:
					keyevent = "[Space]"
					space()
				case termbox.KeyBackspace2, termbox.KeyBackspace:
					keyevent = "[Backspace]"
					backspace()
				default:
					keyevent = "[AddRune]"
					addRune(evt.Ch)
				}
			case termbox.EventError:
				log.Printf("[termboxutil]: Error: Encountered termbox.EventError during Readln loop.")
				err = errors.New("termbox.EventError during Readln loop.")
				break readLoop
			}
			redraw(verb, keyevent)
		}
		if t >= float64(sdur) {
			//newLine() // assumes the user is finished with the word they were working on. would require additional logic to disregard the word during discrepancy check
			break readLoop
		}
	}
	log.Printf("[termboxutil]: Read loop ended:\n\ttime: %.2f\n\twordHistory: %s", t, wordHistory)
	return wordHistory, err
}

func addRune(r rune) {
	s := string(r)
	crntwrd += s
}

func space() {
	if len(crntwrd) > 0 {
		wordHistory = append(wordHistory, crntwrd)
		crntwrd = ""
	}
}

func backspace() {
	if len(crntwrd) == 0 {
		if len(wordHistory) > 0 {
			crntwrd = wordHistory[len(wordHistory)-1]
			wordHistory = wordHistory[:len(wordHistory)-1]
		}
	} else {
		crntwrd = crntwrd[:len(crntwrd)-1]
	}
}
