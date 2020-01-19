package main

import (
	"log"
	"math/rand"
	"strings"

	tbutil "github.com/JosephZoeller/project-0/termboxutil"
	"github.com/JosephZoeller/project-0/timer"
)

var userWords, prgmWords []string

func resetTypeTest() {
	log.Println("[typetest]: Resetting Stopwatch")
	_, running := timer.CheckStopWatch()
	if running {
		timer.PauseStopWatch()
	}
	timer.ResetStopWatch()
	log.Println("[typetest]: Resetting global slices")
	userWords = make([]string, 0)
	prgmWords = make([]string, 0)
}

func loopTestInput(dur int, free, verb bool) float64 {
	tbutil.Write(0, 0, tbutil.COLDEF, tbutil.COLDEF, "Start typing!")
	log.Println("[typetest]: Test started...")
	// buffer required in the event that the user presses the escape key or an error occurs after the display timer is up, but the test hasn't been completed
	cdQuit := make(chan bool, 2)
	defer close(cdQuit)

	go tbutil.CountDown(50, 0, dur, "Time Remaining: %s Seconds...", cdQuit)

	var t float64
	var er error
	for {
		t, _ = timer.CheckStopWatch()
		if t >= float64(dur) { // todo try to remove cast
			break
		}
		func() {
			if !free {
				rndmSnt := sentences[rand.Intn(len(sentences))]
				tbutil.Write(0, 2, tbutil.COLDEF, tbutil.COLDEF, rndmSnt)
				prgmWords = append(prgmWords, strings.Split(rndmSnt, " ")...)
			}
			var u []string
			u, er = tbutil.Readln(dur, verb)
			userWords = append(userWords, u...)
			//tbutil.Write(0, 8, tb.ColorRed, tbutil.COLDEF, fmt.Sprint(userWords))
		}()
		if er != nil {
			log.Println("[typetest]: Exiting main loop, stopping timer goroutine.")
			cdQuit <- true
			break
		}
	}
	log.Println("[typetest]: Test ended.")
	return t
}

//RunTypeTest initiates the typing test for the user, 
// accepting an int for the test duration in seconds,
// a bool for freestyle testing (testing without a writing prompt),
// and a bool for displaying verbose, real-time analytics during the test.
// Returns the user's typed words, the computer's writing prompts and the time spent on the test.
func RunTypeTest(dur int, free, verb *bool) ([]string, []string, float64) {
	resetTypeTest()

	timer.BeginStopWatch()
	t := loopTestInput(dur, *free, *verb)
	timer.PauseStopWatch()

	log.Printf("[typetest]: Elapsed test time: %.2f", t)
	if !*free {
		log.Printf("[typetest]: Computer Generated text: \t%s", prgmWords)
	}
	log.Printf("[typetest]: User Generated text: \t\t%s", userWords)
	return userWords, prgmWords, t
}
