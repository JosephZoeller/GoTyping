package main

import (
	"bufio"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/JosephZoeller/project-0/timer"
	"github.com/nsf/termbox-go"
)

var user *string
var freestyle *bool
var duration *string

const SENTENCEFILE string = "sentences.txt"

var sentences []string

func init() {
	user = flag.String("user", strings.Title(os.Getenv("USER")), "Defaults to the Operating System's current username.")
	freestyle = flag.Bool("free", false, "The user has no writing prompt and can type whatever they please.")
	duration = flag.String("dur", "0:15", "The length of time that the typing test will last. Format as <Minutes>:<Seconds>")
	flag.Parse()

	f, _ := os.Open(SENTENCEFILE)
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		sentences = append(sentences, sc.Text())
	}

	rand.Seed(time.Now().UnixNano())
}

func runSentenceTest() {
	tbCountDown(3, 0, 0, "%s")
	cd, _ := timer.ParseCountDown(*duration)
	timer.BeginStopWatch()
	tbMessage(0, 0, coldef, coldef, "Start typing!")
	var rnd int
	var refsnt string
	var rfwrds, usrwrds, ttlwrds []string
	var cw, lw, incorrect int
	var ttlcw, ttllw int
	go tbCountDown(cd, 50, 0, "Time Remaining: %s Seconds...")
	var t float64
	for {
		t, _ = timer.CheckStopWatch()
		if t >= float64(cd) {
			break
		}
		rnd = rand.Intn(len(sentences)) // loop
		refsnt = sentences[rnd]
		rfwrds = strings.Split(refsnt, " ")
		tbMessage(0, 2, coldef, coldef, refsnt) // show reference sentence
		cw, lw, usrwrds = readLoopSentence(cd)  // compare sentences
		ttlcw += cw
		ttllw += lw
		ttlwrds = append(ttlwrds, rfwrds...)
	}

	incorrect = compare(ttlwrds, usrwrds)
	timer.PauseStopWatch()
	t, _ = timer.CheckStopWatch()
	tbprintStats(ttlcw, ttllw, t) // check acccuracy
	tbprintAccur(incorrect, len(ttlwrds))
	keyContinue(true)
}

func compare(src, usr []string) int {
	wrong := 0 // src has the full sentence, but the user might not have typed the entire sentence  before the time was up

	// maybe upon finding one wrong, check the usr[i + 1], and then for the next word check [i - 1] to see if they got back on track
	for i, sr := range src {
		if i < len(usr) && sr != usr[i] {
			wrong++
		}
	}
	return wrong
}

func tbprintAccur(wrng, ttl int) {

	tbMessage(60, 2, termbox.ColorBlue, coldef, fmt.Sprintf("Total words: %d", ttl))
	tbMessage(60, 3, termbox.ColorBlue, coldef, fmt.Sprintf("Words missed: %d", wrng))
	tbMessage(60, 4, termbox.ColorBlue, coldef, fmt.Sprintf("Accuracy: %% %.2f", float64(ttl-wrng)/float64(ttl)*100))
}

func runTest() {
	tbCountDown(3, 0, 0, "%s")
	cd, _ := timer.ParseCountDown(*duration)
	timer.BeginStopWatch()
	tbMessage(0, 0, coldef, coldef, "Start typing!")
	go tbCountDown(cd, 50, 0, "Time Remaining: %s Seconds...")
	cw, lw := readLoop(cd)
	timer.PauseStopWatch()
	t, _ := timer.CheckStopWatch()
	tbprintStats(cw, lw, t) // get count
	keyContinue(true)
}

func tbprintStats(c, l int, t float64) {
	cfl := float64(c)
	lfl := float64(l)
	termbox.Clear(coldef, coldef)
	termbox.HideCursor()
	tbMessage(0, 0, termbox.ColorBlue, coldef, fmt.Sprintf("Seconds to complete: %.2f", t))

	tbMessage(0, 2, termbox.ColorBlue, coldef, fmt.Sprintf("Words written: %d", c))
	tbMessage(0, 3, termbox.ColorBlue, coldef, fmt.Sprintf("Words per second: %.2f", cfl/t))
	tbMessage(0, 4, termbox.ColorBlue, coldef, fmt.Sprintf("Words per minute: %.2f", cfl/t*60))

	tbMessage(27, 2, termbox.ColorBlue, coldef, fmt.Sprintf("Characters written: %d", l))
	tbMessage(27, 3, termbox.ColorBlue, coldef, fmt.Sprintf("Characters per second: %.2f", lfl/t))
	tbMessage(27, 4, termbox.ColorBlue, coldef, fmt.Sprintf("Characters per minute: %.2f", lfl/t*60))

	tbMessage(0, 6, coldef, coldef, "Press the enter key to end the program...")
}
