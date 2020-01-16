package main

import (
	"fmt"

	tbutil "github.com/JosephZoeller/project-0/termboxutil"
	tb "github.com/nsf/termbox-go"
)

func getDiscrepancyCount(userWords, prgmWords []string) int {
	wrong := 0

	// maybe upon finding one wrong, check the usr[i + 1], and then for the next word check [i - 1] to see if they got back on track
	for i, prgmWord := range prgmWords {
		if i < len(userWords) && prgmWord != userWords[i] {
			wrong++
		}
	}
	return wrong
}

func getByteCount(strslice []string) int {
	count := 0
	for _, str := range strslice {
		count += len(str)
	}
	return count
}
func tbprintAccur(wrng, ttl int) {
	tbutil.Write(60, 2, tb.ColorBlue, tbutil.COLDEF, fmt.Sprintf("Total words: %d", ttl))
	tbutil.Write(60, 3, tb.ColorBlue, tbutil.COLDEF, fmt.Sprintf("Words missed: %d", wrng))
	tbutil.Write(60, 4, tb.ColorBlue, tbutil.COLDEF, fmt.Sprintf("Accuracy: %% %.2f", float64(ttl-wrng)/float64(ttl)*100))
}

func tbprintStats(c, l int, t float64) {
	cfl := float64(c)
	lfl := float64(l)
	tb.Clear(tbutil.COLDEF, tbutil.COLDEF)
	tb.HideCursor()
	tbutil.Write(0, 0, tb.ColorBlue, tbutil.COLDEF, fmt.Sprintf("Seconds to complete: %.2f", t))

	tbutil.Write(0, 2, tb.ColorBlue, tbutil.COLDEF, fmt.Sprintf("Words written: %d", c))
	tbutil.Write(0, 3, tb.ColorBlue, tbutil.COLDEF, fmt.Sprintf("Words per second: %.2f", cfl/t))
	tbutil.Write(0, 4, tb.ColorBlue, tbutil.COLDEF, fmt.Sprintf("Words per minute: %.2f", cfl/t*60))

	tbutil.Write(27, 2, tb.ColorBlue, tbutil.COLDEF, fmt.Sprintf("Characters written: %d", l))
	tbutil.Write(27, 3, tb.ColorBlue, tbutil.COLDEF, fmt.Sprintf("Characters per second: %.2f", lfl/t))
	tbutil.Write(27, 4, tb.ColorBlue, tbutil.COLDEF, fmt.Sprintf("Characters per minute: %.2f", lfl/t*60))

	tbutil.Write(0, 6, tbutil.COLDEF, tbutil.COLDEF, "Press the enter key to end the program...")
}
