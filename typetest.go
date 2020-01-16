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
	duration = flag.String("dur", "0:30", "The length of time that the typing test will last. Format as <Minutes>:<Seconds>")
	flag.Parse()

	f, _ := os.Open(SENTENCEFILE)
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		sentences = append(sentences, sc.Text())
	}

	rand.Seed(time.Now().UnixNano())
}

var userWords, prgmWords []string

func loopTestInput(dur int, free bool) float64 {
	tbCountDown(0, 0, 3, "%s")
	timer.ResetStopWatch()
	tbMessage(0, 0, COLDEF, COLDEF, "Start typing!")
	timer.BeginStopWatch()
	go tbCountDown(50, 0, dur, "Time Remaining: %s Seconds...")
	var t float64
	for {
		t, _ = timer.CheckStopWatch()
		if t >= float64(dur) { // todo try to remove cast
			break
		}
		func() {
			if !free {
				rndmSnt := sentences[rand.Intn(len(sentences))]
				tbMessage(0, 2, COLDEF, COLDEF, rndmSnt)
				prgmWords = append(prgmWords, strings.Split(rndmSnt, " ")...)
			}
			userWords = append(userWords, readln(dur)...)
			tbMessage(0, 7, termbox.ColorRed, COLDEF, fmt.Sprint(userWords))
		}()
	}
	timer.PauseStopWatch()
	return t
}

func resetTypeTest() {
	userWords = make([]string, 0)
	prgmWords = make([]string, 0)
}

func runTypeTest(dur *string, free *bool) ([]string, []string, float64) {
	cd, _ := timer.ParseCountDown(*dur)
	resetTypeTest()
	t := loopTestInput(cd, *free)
	return userWords, prgmWords, t
}

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
	tbMessage(60, 2, termbox.ColorBlue, COLDEF, fmt.Sprintf("Total words: %d", ttl))
	tbMessage(60, 3, termbox.ColorBlue, COLDEF, fmt.Sprintf("Words missed: %d", wrng))
	tbMessage(60, 4, termbox.ColorBlue, COLDEF, fmt.Sprintf("Accuracy: %% %.2f", float64(ttl-wrng)/float64(ttl)*100))
}

func tbprintStats(c, l int, t float64) {
	cfl := float64(c)
	lfl := float64(l)
	termbox.Clear(COLDEF, COLDEF)
	termbox.HideCursor()
	tbMessage(0, 0, termbox.ColorBlue, COLDEF, fmt.Sprintf("Seconds to complete: %.2f", t))

	tbMessage(0, 2, termbox.ColorBlue, COLDEF, fmt.Sprintf("Words written: %d", c))
	tbMessage(0, 3, termbox.ColorBlue, COLDEF, fmt.Sprintf("Words per second: %.2f", cfl/t))
	tbMessage(0, 4, termbox.ColorBlue, COLDEF, fmt.Sprintf("Words per minute: %.2f", cfl/t*60))

	tbMessage(27, 2, termbox.ColorBlue, COLDEF, fmt.Sprintf("Characters written: %d", l))
	tbMessage(27, 3, termbox.ColorBlue, COLDEF, fmt.Sprintf("Characters per second: %.2f", lfl/t))
	tbMessage(27, 4, termbox.ColorBlue, COLDEF, fmt.Sprintf("Characters per minute: %.2f", lfl/t*60))

	tbMessage(0, 6, COLDEF, COLDEF, "Press the enter key to end the program...")
}
