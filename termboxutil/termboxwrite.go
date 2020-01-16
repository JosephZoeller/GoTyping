package termboxutil

import (
	"fmt"
	"strconv"
	"time"

	"github.com/nsf/termbox-go"
)

const COLDEF termbox.Attribute = termbox.ColorDefault

func Write(xStart int, yStart int, foreColor termbox.Attribute, backColor termbox.Attribute, message string) (int, int) {
	x := xStart
	y := yStart
	stdWidth, _ := termbox.Size()

	for _, ch := range message {
		termbox.SetCell(x, y, ch, foreColor, backColor)
		if x > stdWidth && ch == ' ' {
			y++
			x = xStart
		} else {
			x++
		}
	}

	for j := x; j < stdWidth; j++ {
		termbox.SetCell(j, y, ' ', foreColor, backColor)
	}

	termbox.Flush()
	return x, y
}

func redraw() {
	sntX, sntY := Write(0, 1, COLDEF, COLDEF, snt)
	termbox.SetCursor(sntX, sntY)

	Write(0, 3, COLDEF, COLDEF, ("Event: " + keyevent))
	Write(0, 4, COLDEF, COLDEF, fmt.Sprintf("Word Bank: %s", wordHistory))
	Write(0, 5, COLDEF, COLDEF, "Current word: "+crntwrd)
}

func CountDown(x, y, cd int, frmt string) {
	for cd > 0 {
		Write(x, y, COLDEF, COLDEF, fmt.Sprintf(frmt, strconv.Itoa(cd)))
		time.Sleep(time.Second)
		cd--
	}
	Write(x, y, COLDEF, COLDEF, fmt.Sprintf(frmt, "0"))
}
