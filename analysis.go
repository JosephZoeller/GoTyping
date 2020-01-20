package main

import (
	"fmt"
	tbutil "github.com/JosephZoeller/project-0/termboxutil"
	tb "github.com/nsf/termbox-go"
	"log"
)

func getDiscrepancyCount(userWords, prgmWords []string) int {
	wrong := 0
	for i, prgmWord := range prgmWords {
		if i < len(userWords) && prgmWord != userWords[i] {
			wrong++
		}
	}
	log.Printf("[analysis]: Total typing discrepancies found: %d", wrong)
	return wrong
}

func getByteCount(strslice []string) int {
	count := 0
	for _, str := range strslice {
		count += len(str)
	}
	log.Printf("[analysis]: Measuring byte count:\n\tString: %s\n\tByte count: %d", strslice, count)
	return count
}
func tbprintAccur(wrng, ttl int, cheat *bool) {
	if *cheat {
		wrng /= 3
		log.Printf("[analysis]: Cheat mode enabled: Displaying %d (# wrong answers / 3)", wrng)
	}
	tbutil.Write(60, 2, tb.ColorBlue, tbutil.COLDEF, fmt.Sprintf("Total words: %d", ttl))
	log.Printf("[analysis]: Total words: %d", ttl)
	tbutil.Write(60, 3, tb.ColorBlue, tbutil.COLDEF, fmt.Sprintf("Words missed: %d", wrng))
	log.Printf("[analysis]: Words missed: %d", wrng)
	tbutil.Write(60, 4, tb.ColorBlue, tbutil.COLDEF, fmt.Sprintf("Accuracy: %% %.2f", float64(ttl-wrng)/float64(ttl)*100))
	log.Printf("[analysis]: Accuracy: %% %.2f", float64(ttl-wrng)/float64(ttl)*100)
}

func tbprintStats(wrds []string, t float64, cheat *bool) {
	wordLen := len(wrds)
	wordLenFl := float64(wordLen)
	charLen := getByteCount(wrds)
	charLenFl := float64(charLen)

	tb.Clear(tbutil.COLDEF, tbutil.COLDEF)
	tb.HideCursor()
	tbutil.Write(0, 0, tb.ColorBlue, tbutil.COLDEF, fmt.Sprintf("Seconds to complete: %.2f", t))

	if *cheat {
		t = t / 3
		log.Printf("[analysis]: Cheat mode enabled: Calculating for time %.2f (test duration / 3)", t)
	}

	tbutil.Write(0, 2, tb.ColorBlue, tbutil.COLDEF, fmt.Sprintf("Words written: %d", wordLen))
	log.Printf("[analysis]: Words written: %d", charLen)
	tbutil.Write(0, 3, tb.ColorBlue, tbutil.COLDEF, fmt.Sprintf("Words per second: %.2f", wordLenFl/t))
	log.Printf("[analysis]: Words per second: %.2f", wordLenFl/t)
	tbutil.Write(0, 4, tb.ColorBlue, tbutil.COLDEF, fmt.Sprintf("Words per minute: %.2f", wordLenFl/t*60))
	log.Printf("[analysis]: Words per minute: %.2f", wordLenFl/t*60)

	tbutil.Write(27, 2, tb.ColorBlue, tbutil.COLDEF, fmt.Sprintf("Characters written: %d", charLen))
	log.Printf("[analysis]: Characters written: %d", charLen)
	tbutil.Write(27, 3, tb.ColorBlue, tbutil.COLDEF, fmt.Sprintf("Characters per second: %.2f", charLenFl/t))
	log.Printf("[analysis]: Characters per second: %.2f", charLenFl/t)
	tbutil.Write(27, 4, tb.ColorBlue, tbutil.COLDEF, fmt.Sprintf("Characters per minute: %.2f", charLenFl/t*60))
	log.Printf("[analysis]: Characters per minute: %.2f", charLenFl/t*60)

	tbutil.Write(0, 6, tbutil.COLDEF, tbutil.COLDEF, "Press the enter key to end the program...")
}
