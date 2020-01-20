package main

import (
	"bufio"
	"log"
	"os"

	"github.com/JosephZoeller/project-0/cmd/phase"
	tbutil "github.com/JosephZoeller/project-0/pkg/termboxutil"
	"github.com/JosephZoeller/project-0/pkg/timer"
	"github.com/nsf/termbox-go"
)

func main() {
	file, logcrerr := os.Create(LOGFILE)
	if logcrerr != nil {
		log.Fatal(logcrerr.Error())
	}
	log.SetOutput(file)
	log.Println("[main]: Program Start...")
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
		if len(sentences) < 2 {
			sentences = append(sentences, "There are not enough sentences in the sentences.txt file.", "This is a default sentence, since sentences.txt was found lacking.")
		}
	}

	dur, durerr := timer.ParseCountDown(*duration)
	if durerr != nil {
		log.Printf("[main]: An error occurred while parsing the custom duration \"%s\". Defaulting to 0:30.", *duration)
		tbutil.Write(0, 0, tbutil.COLDEF, tbutil.COLDEF, "Error while parsing duration. Defaulting to 0:30")
		dur = 30
	}

	// test briefing
	log.Println("[main]: Displaying preface...")
	phase.ShowPreface(*user, *freestyle)

	// test
	log.Printf("[main]: Beginning test... [User: %s, Duration: %d, Freestyle: %t, Verbose: %t, Cheat: %t]", *user, dur, *freestyle, *debug, *cheat)
	uwrds, wrngCnt, t := phase.RunTypeTest(dur, freestyle, debug, sentences)
	log.Printf("[main]: Test Complete...")

	// test debriefing
	log.Println("[main]: Displaying speed statistics...")
	phase.TbprintStats(uwrds, t, cheat)
	if !*freestyle {
		log.Println("[main]: Displaying accuracy statistics...")
		phase.TbprintAccur(len(uwrds), wrngCnt, t, cheat)
	}
	tbutil.KeyContinue(true)
	log.Println("[main]: Exiting Program...")
}