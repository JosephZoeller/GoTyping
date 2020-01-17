package main

import (
	"flag"
	"math/rand"
	"os"
	"strings"
	"time"
)

var user *string
var freestyle, verbose, cheat *bool
var duration *string

const SENTENCEFILE string = "sentences.txt"
const LOGFILE string = "log.txt"
var sentences []string

func init() {
	user = flag.String("user", strings.Title(os.Getenv("USER")), "Defaults to the Operating System's current username.")
	freestyle = flag.Bool("free", false, "Removes the writing prompt. The user can type without restriction and accuracy won't be measured.")
	verbose = flag.Bool("v", false, "Displays under-the-hood details during the test.")
	cheat = flag.Bool("wingman", false, "Fudges the test results to impress your peers.")
	duration = flag.String("dur", "0:30", "The length of time that the typing test will last. Format as <Minutes>:<Seconds>")
	flag.Parse()
	rand.Seed(time.Now().UnixNano())
}
