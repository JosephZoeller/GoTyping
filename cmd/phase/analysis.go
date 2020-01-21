// Package phase project-0 Typing Test for 200106-uta-go
package phase

import (
	"fmt"
	tbutil "github.com/JosephZoeller/project-0/pkg/termboxutil"
	tb "github.com/nsf/termbox-go"
	"log"
)

// getDiscrepancyCount compares the user-generated words against the program-generated words.
// accepts both user and computer text as []string and returns the number of mismatches found.
// If, hypothetically, a user skips or adds a word, the strings would be offset by one when comparing the two inputs.
// This would lead to a cascade of discrepancies, so this function is considered fragile in that regard.
func getDiscrepancyCount(userWords, prgmWords []string) int {
	wrong := 0

	// maybe upon finding one wrong, check the usr[i + 1], and then for the next word check [i - 1] to see if they got back on track
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

// TbprintAccur displays and logs accuracy calculations, based on the data accrued during the test.
// Accepts ints for the user's total words typed and those that they mistyped,
// as well as a float64 representing the test's elapsed time in seconds
// and the cheat boolean, which thirds the errors accrued and time elapsed.
func TbprintAccur(ttl, wrng int, t float64, cheat *bool) {
	if *cheat {
		wrng /= 3
		t /= 3
		log.Printf("[analysis]: Cheat mode enabled: Displaying %d (# wrong answers / 3)", wrng)
	}
	crrctFL := float64(ttl - wrng)
	ttlFl := float64(ttl)
	tbutil.Write(60, 2, tbutil.COLDEF, tbutil.COLDEF, fmt.Sprintf("Words missed: %d", wrng))
	log.Printf("[analysis]: Words missed: %d", wrng)
	tbutil.Write(60, 3, tbutil.COLDEF, tbutil.COLDEF, fmt.Sprintf("Accuracy: %% %.2f", crrctFL/ttlFl*100))
	log.Printf("[analysis]: Accuracy: %% %.2f", crrctFL/ttlFl*100)
	tbutil.Write(60, 4, tbutil.COLDEF, tbutil.COLDEF, fmt.Sprintf("Adjusted words per minute: %.2f", crrctFL/t*60))
	log.Printf("[analysis]: Adjusted words per minute: %.2f", crrctFL/t*60)
}

// TbprintStats displays and logs general calculations, based on the data accrued during the test.
// Accepts a slice of words typed by the user, as well as a float64 representing the test's elapsed time in seconds
// and the cheat boolean, which thirds the time elapsed.
func TbprintStats(wrds []string, t float64, cheat *bool) {
	wordLen := len(wrds)
	wordLenFl := float64(wordLen)
	charLen := getByteCount(wrds)
	charLenFl := float64(charLen)

	tb.Clear(tbutil.COLDEF, tbutil.COLDEF)
	tb.HideCursor()
	tbutil.Write(0, 0, tbutil.COLDEF, tbutil.COLDEF, fmt.Sprintf("Seconds to complete: %.2f", t))

	if *cheat {
		t /= 3
		log.Printf("[analysis]: Cheat mode enabled: Calculating for time %.2f (test duration / 3)", t)
	}

	tbutil.Write(0, 2, tbutil.COLDEF, tbutil.COLDEF, fmt.Sprintf("Words written: %d", wordLen))
	log.Printf("[analysis]: Words written: %d", charLen)
	tbutil.Write(0, 3, tbutil.COLDEF, tbutil.COLDEF, fmt.Sprintf("Words per second: %.2f", wordLenFl/t))
	log.Printf("[analysis]: Words per second: %.2f", wordLenFl/t)
	tbutil.Write(0, 4, tbutil.COLDEF, tbutil.COLDEF, fmt.Sprintf("Words per minute: %.2f", wordLenFl/t*60))
	log.Printf("[analysis]: Words per minute: %.2f", wordLenFl/t*60)

	tbutil.Write(27, 2, tbutil.COLDEF, tbutil.COLDEF, fmt.Sprintf("Characters written: %d", charLen))
	log.Printf("[analysis]: Characters written: %d", charLen)
	tbutil.Write(27, 3, tbutil.COLDEF, tbutil.COLDEF, fmt.Sprintf("Characters per second: %.2f", charLenFl/t))
	log.Printf("[analysis]: Characters per second: %.2f", charLenFl/t)
	tbutil.Write(27, 4, tbutil.COLDEF, tbutil.COLDEF, fmt.Sprintf("Characters per minute: %.2f", charLenFl/t*60))
	log.Printf("[analysis]: Characters per minute: %.2f", charLenFl/t*60)

	tbutil.Write(0, 6, tbutil.COLDEF, tbutil.COLDEF, "Press the enter key to end the program...")
}
