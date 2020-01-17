package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"

	tbutil "github.com/JosephZoeller/project-0/termboxutil"
	"github.com/nsf/termbox-go"
)

func main() {
	log.Println("[main]: Program Start...")
	file, logcrerr := os.Create(LOGFILE)
	if logcrerr != nil {
		log.Fatal(logcrerr.Error())
	}
	log.SetOutput(file)
	defer file.Close()

	err := termbox.Init()
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer termbox.Close()
	termbox.SetInputMode(termbox.InputEsc)

	f, senfilerr := os.Open(SENTENCEFILE)
	if senfilerr != nil {
		log.Println("[main]: " + SENTENCEFILE + " could not be opened. Defaulting to freestyle mode. Error: " + senfilerr.Error())
		*freestyle = true
	} else {
		sc := bufio.NewScanner(f)
		for sc.Scan() {
			sentences = append(sentences, sc.Text())
		}
	}
	dur, durerr := ParseCountDown(*duration)
	if durerr != nil {
		log.Printf("[main]: An error occurred while parsing the custom duration \"%s\". Defaulting to 0:30.", *duration)
		tbutil.Write(0, 0, tbutil.COLDEF, tbutil.COLDEF, "Error while parsing duration. Defaulting to 0:30")
		dur = 30
	}

	// test briefing
	log.Println("[main]: Displaying preface...")
	showPreface(*user)

	// test
	log.Printf("[main]: Beginning test... [User: %s, Duration: %d, Freestyle: %t, Verbose: %t, Cheat: %t]", *user, dur, *freestyle, *verbose, *cheat)
	uwrds, pgwrds, t := runTypeTest(dur, freestyle, verbose)
	log.Printf("[main]: Test Complete...")
	er := getDiscrepancyCount(uwrds, pgwrds)

	// test debriefing
	log.Println("[main]: Displaying speed statistics...")
	tbprintStats(len(uwrds), getByteCount(uwrds), t, cheat)
	if len(prgmWords) > 0 {
		log.Println("[main]: Displaying accuracy statistics...")
		tbprintAccur(er, len(uwrds), cheat)
	}
	tbutil.KeyContinue(true)
	log.Println("[main]: Exiting Program...")
}

func ParseCountDown(a string) (int, error) {
	s := 0
	aDlm := strings.Split(a, ":")
	for i, a := range aDlm {
		t, e := strconv.Atoi(a)
		if e != nil {
			return -1, e
		}
		switch i {
		case 0:
			s += (t * 60)
		case 1:
			s += t
		}
	}
	return s, nil
}
