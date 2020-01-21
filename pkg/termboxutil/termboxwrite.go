// Package termboxutil project-0 Typing Test for 200106-uta-go
package termboxutil

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"time"

	tb "github.com/nsf/termbox-go"
)

const COLDEF tb.Attribute = tb.ColorDefault

// Write displays the input string on the terminal through termbox.
// Accepts an x, y coordinate, a foreground and background color, and a message to print.
// returns the coordinates of the end of the message string.
// Note: Write will draw over previously written text, both on and to the right of the x coordinate.
// In other words, in order for two or more strings to be displayed on the same line,
// They must be written in order from left to right,
func Write(xStart int, yStart int, foreColor tb.Attribute, backColor tb.Attribute, message string) (int, int) {
	x := xStart
	y := yStart
	stdWidth, _ := tb.Size()
	if stdWidth >= 10 {
		stdWidth -= 10
	}

	for _, ch := range message {
		tb.SetCell(x, y, ch, foreColor, backColor)
		if x > stdWidth && ch == ' ' {
			y++
			x = xStart
		} else {
			x++
		}
	}

	for j := x; j < stdWidth+10; j++ {
		tb.SetCell(j, y, ' ', foreColor, backColor)
	}

	tb.Flush()
	return x, y
}

// redraw displays the line of text that the user is currently inputting. For each space, added rune or backspace, the input is redrawn on the screen.
// accepts a boolean to turn on verbose (debug) information regarding the user input, and a string which displays what triggered the redraw.
func redraw(verb bool, keyEvent string) {
	var snt string
	for _, wrd := range wordHistory {
		snt += wrd + " "
	}
	snt += crntwrd
	sntX, sntY := Write(0, 3, COLDEF, COLDEF, snt)
	if verb {
		drawRTStats(keyEvent)
	}
	tb.SetCursor(sntX, sntY)
	tb.Flush() // otherwise SetCursor will need to wait for the next redraw to move which is nauseating
}
// drawRTStats draws the real-time stats to the terminal.
// includes the trigger event, the user's word history and the current word that the user is working on.
func drawRTStats(keyEvent string) {
	Write(0, 5, tb.ColorGreen, COLDEF, ("Event: " + keyEvent))
	//t, _ := timer.CheckStopWatch()
	//Write(50, 5, tb.ColorGreen, COLDEF, fmt.Sprintf("Average Speed: %.2f WPM", float64(len(wordHistory))/t*60)) // Inaccurate after the first readLn loop since wordHistory is reset, but timer is not.
	Write(0, 6, tb.ColorGreen, COLDEF, fmt.Sprintf("Word Bank: %s", wordHistory))
	Write(0, 7, tb.ColorGreen, COLDEF, "Current word: "+crntwrd)
}

// CountDown displays a countdown to the terminal, updating every second until it reaches 0.
// accepts an x,y starting coordinate, a formatted string with which the countdown integer will be displayed as,
// and an abort channel, which will kill the function when it recieves 'true' from the channel. 
func CountDown(x, y, cd int, frmt string, abort chan bool) error {
	if cd < 1 {
		log.Printf("[termboxutil]: Countdown value must be greater than 0 (input: %d)", cd)
		return errors.New(fmt.Sprintf("[termboxutil]: Countdown value must be greater than 0 (input: %d)", cd))
	}
	log.Printf("[termboxutil]: Countdown initiated: Coordinates (%d,%d), %d Seconds", x, y, cd)
	col := COLDEF
	for cd > 0 {
		select {
		case <-abort:
			log.Printf("[termboxutil]: Countdown aborted: Coordinates (%d, %d), %d Seconds", x, y, cd)
			return (errors.New(fmt.Sprintf("[termboxutil]: Countdown aborted: Coordinates (%d, %d), %d Seconds", x, y, cd)))
		default:
			if cd <= 10 {
				col = tb.ColorRed
			}
			Write(x, y, col, COLDEF, fmt.Sprintf(frmt, strconv.Itoa(cd)))
			time.Sleep(time.Second)
			cd--
		}
	}
	Write(x, y, col, COLDEF, fmt.Sprintf(frmt, "0"))
	log.Printf("[termboxutil]: Countdown completed: Coordinates (%d,%d)", x, y)
	return nil
}
