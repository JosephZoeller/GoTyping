package main

import (
	"fmt"
	"strings"
	"math/rand"

	tbutil "github.com/JosephZoeller/project-0/termboxutil"
	"github.com/JosephZoeller/project-0/timer"
	tb "github.com/nsf/termbox-go"
)

var userWords, prgmWords []string

func resetTypeTest() {
	userWords = make([]string, 0)
	prgmWords = make([]string, 0)
}

func loopTestInput(dur int, free bool) float64 {
	timer.ResetStopWatch()
	tbutil.Write(0, 0, tbutil.COLDEF, tbutil.COLDEF, "Start typing!")
	timer.BeginStopWatch()
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
			userWords = append(userWords, tbutil.Readln(dur)...)
			tbutil.Write(0, 7, tb.ColorRed, tbutil.COLDEF, fmt.Sprint(userWords))
		}()
	}
	timer.PauseStopWatch()
	return t
}

func runTypeTest(dur int, free *bool) ([]string, []string, float64) {
	resetTypeTest()
	t := loopTestInput(dur, *free)
	return userWords, prgmWords, t
}