package termboxutil

import (
	"errors"
	"fmt"
	"github.com/JosephZoeller/project-0/timer"
	tb "github.com/nsf/termbox-go"
	"log"
	"strconv"
	"time"
)

const COLDEF tb.Attribute = tb.ColorDefault

func Write(xStart int, yStart int, foreColor tb.Attribute, backColor tb.Attribute, message string) (int, int) {
	x := xStart
	y := yStart
	stdWidth, _ := tb.Size()

	for _, ch := range message {
		tb.SetCell(x, y, ch, foreColor, backColor)
		if x > stdWidth && ch == ' ' {
			y++
			x = xStart
		} else {
			x++
		}
	}

	for j := x; j < stdWidth; j++ {
		tb.SetCell(j, y, ' ', foreColor, backColor)
	}

	tb.Flush()
	return x, y
}

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

func drawRTStats(keyEvent string) { // Assumes timer
	Write(0, 5, tb.ColorGreen, COLDEF, ("Event: " + keyEvent))
	t, _ := timer.CheckStopWatch()
	Write(50, 5, tb.ColorGreen, COLDEF, fmt.Sprintf("Average Speed: %.2f WPM", float64(len(wordHistory))/t*60))
	Write(0, 6, tb.ColorGreen, COLDEF, fmt.Sprintf("Word Bank: %s", wordHistory))
	Write(0, 7, tb.ColorGreen, COLDEF, "Current word: "+crntwrd)
}

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
