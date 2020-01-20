package main

import (
	"errors"
	"log"
	"math/rand"
	"strings"

	tbutil "github.com/JosephZoeller/project-0/termboxutil"
	"github.com/JosephZoeller/project-0/timer"
)

func loopTestInput(dur int, free, verb bool) ([]string, int) {
	tbutil.Write(0, 0, tbutil.COLDEF, tbutil.COLDEF, "Start typing!")
	log.Println("[typetest]: Test started...")
	cdQuit := make(chan bool, 2) // buffer required in the event that the user presses the escape key or an error occurs after the display timer is up, but the test hasn't been completed
	defer close(cdQuit)

	go tbutil.CountDown(50, 0, dur, "Time Remaining: %s Seconds...", cdQuit)

	var er error
	var rndmSnt string
	userWords := make([]string, 0)
	wrngCnt := 0
	t, _ := timer.CheckStopWatch()
	for {
		if er != nil || t >= float64(dur) {
			log.Println("[typetest]: Exiting main loop, stopping timer goroutine.")
			cdQuit <- true
			break
		}

		func() {
			if !free {
				rndmSnt, _ = getRandomSentencePsuedo(rndmSnt)
				tbutil.Write(0, 2, tbutil.COLDEF, tbutil.COLDEF, rndmSnt)
				log.Printf("[typetest]: Computer Generated text: \t%s", rndmSnt)
			}
			var u []string
			u, er = tbutil.Readln(dur, verb)
			log.Printf("[typetest]: User Generated text: \t\t%s", u)

			userWords = append(userWords, u...)
			if !free {
				wrngCnt += getDiscrepancyCount(u, strings.Split(rndmSnt, " "))
			}
		}()

		t, _ = timer.CheckStopWatch()
	}
	log.Println("[typetest]: Test ended.")
	return userWords, wrngCnt
}

func getRandomSentencePsuedo(lastsnt string) (string, error) {
	if len(sentences) > 1 {
		for {
			rndmSnt := sentences[rand.Intn(len(sentences))]
			if rndmSnt != lastsnt {
				return rndmSnt, nil
			}
		}
	}
	return "", errors.New("Insufficient number of writing prompts to choose from (minimum 2).")
}

// RunTypeTest initiates the typing test for the user,
// accepting an int for the test duration in seconds,
// a bool for freestyle testing (testing without a writing prompt),
// and a bool for displaying verbose, real-time analytics during the test.
// Returns the user's typed words, the computer's writing prompts and the time spent on the test.
func RunTypeTest(dur int, free, verb *bool) ([]string, int, float64) {
	timer.PrimeStopWatch()

	timer.BeginStopWatch()
	userWords, wrngCnt := loopTestInput(dur, *free, *verb)
	timer.PauseStopWatch()
	t, _ := timer.CheckStopWatch()

	log.Printf("[typetest]: Elapsed test time: %.2f", t)
	return userWords, wrngCnt, t
}
