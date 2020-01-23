// Package main project-0 Typing Test for 200106-uta-go
package main

import (
	"log"

	"github.com/JosephZoeller/project-0/cmd/phase"
	tbutil "github.com/JosephZoeller/project-0/pkg/termboxutil"
	"github.com/nsf/termbox-go"
)

func main() {

	config()
	defer LogFile.Close()
	defer termbox.Close()

	// briefing
	log.Println("[main]: Displaying preface...")
	phase.ShowPreface(*user, duration, *freestyle)

	// user input
	log.Printf("[main]: Beginning test... [User: %s, Duration: %d, Freestyle: %t, Verbose: %t, Cheat: %t]", *user, duration, *freestyle, *debug, *cheat)
	uwrds, wrngCnt, t := phase.RunTypeTest(duration, freestyle, debug, sentences)
	log.Printf("[main]: Test Complete...")

	// debriefing
	log.Println("[main]: Displaying speed statistics...")
	phase.TbprintStats(uwrds, t, cheat)
	if !*freestyle {
		log.Println("[main]: Displaying accuracy statistics...")
		phase.TbprintAccur(len(uwrds), wrngCnt, t, cheat)
	}
	tbutil.KeyContinue(true)
	log.Println("[main]: Exiting Program...")
}