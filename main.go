package main

import (
	"github.com/nsf/termbox-go"
)

func main() {
	// termbox init
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()
	termbox.SetInputMode(termbox.InputEsc)

	// test briefing
	showPreface()
	
	// test
	uwrds, pgwrds, t := runTypeTest(duration, freestyle)
	er := getDiscrepancyCount(uwrds, pgwrds)

	// test debriefing
	tbprintStats(len(uwrds), getByteCount(uwrds), t)
	if (len(prgmWords) > 0) {
		tbprintAccur(er, len(uwrds))
	}
	keyContinue(true)
}