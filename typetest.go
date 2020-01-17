package main

import (
	"math/rand"
	"strings"
	"log"

	tbutil "github.com/JosephZoeller/project-0/termboxutil"
	"github.com/JosephZoeller/project-0/timer"
)

var userWords, prgmWords []string

func resetTypeTest() {
	log.Println("[typetest]: Resetting global slices")
	userWords = make([]string, 0)
	prgmWords = make([]string, 0)
}

func loopTestInput(dur int, free, verb bool) float64 {
	timer.ResetStopWatch()
	tbutil.Write(0, 0, tbutil.COLDEF, tbutil.COLDEF, "Start typing!")
	timer.BeginStopWatch()
	log.Println("[typetest]: Test started...")
	go tbutil.CountDown(50, 0, dur, "Time Remaining: %s Seconds...")
	var t float64
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
			userWords = append(userWords, tbutil.Readln(dur, verb)...)
			//tbutil.Write(0, 7, tb.ColorRed, tbutil.COLDEF, fmt.Sprint(userWords))
		}()
	}
	timer.PauseStopWatch()
	log.Println("[typetest]: Test ended.")
	return t
}

func runTypeTest(dur int, free, verb *bool) ([]string, []string, float64) {
	resetTypeTest()
	t := loopTestInput(dur, *free, *verb)
	log.Printf("[typetest]: Elapsed test time: %.2f", t)
	if (!*free) {
		log.Printf("[typetest]: Computer Generated text: %s", prgmWords)
	}
	log.Printf("[typetest]: User Generated text: %s", userWords)
	return userWords, prgmWords, t
}
