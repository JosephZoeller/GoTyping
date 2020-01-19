package main

import (
	"flag"
	"math/rand"
	"os"
	"strings"
	"time"
)

var user *string
var freestyle, debug, cheat *bool
var duration *string

const SENTENCEFILE string = "sentences.txt"
const LOGFILE string = "log.txt"

var sentences []string

func init() {
	user = flag.String("u", strings.Title(os.Getenv("USER")), "User - Defaults to the Operating System's current username.")
	freestyle = flag.Bool("f", false, "Freestyle - Removes the writing prompt. The user can type without restriction and accuracy won't be measured.")
	debug = flag.Bool("debug", false, "Debug - Displays under-the-hood details during the test.")
	cheat = flag.Bool("c", false, "Cheat - Fudges the test results to impress your peers.")
	duration = flag.String("d", "0:30", "Duration - The length of time that the typing test will last. Format as <Minutes>:<Seconds>")
	flag.Parse()
	rand.Seed(time.Now().UnixNano())
}
