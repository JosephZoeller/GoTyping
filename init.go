package main

import (
	"bufio"
	"errors"
	"flag"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

var user *string
var freestyle *bool
var duration int

const SENTENCEFILE string = "sentences.txt"

var sentences []string

func init() {

	user = flag.String("user", strings.Title(os.Getenv("USER")), "Defaults to the Operating System's current username.")
	freestyle = flag.Bool("free", false, "The user has no writing prompt and can type whatever they please.")
	var durerr error
	duration, durerr = ParseCountDown(*flag.String("dur", "0:30", "The length of time that the typing test will last. Format as <Minutes>:<Seconds>"))
	flag.Parse()

	if durerr != nil {
		// default
	}
	f, _ := os.Open(SENTENCEFILE)
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		sentences = append(sentences, sc.Text())
	}

	rand.Seed(time.Now().UnixNano())
}

func ParseCountDown(a string) (int, error) {
	s := 0
	aDlm := strings.Split(a, ":")
	for i, a := range aDlm {
		t, _ := strconv.Atoi(a)
		switch i {
		case 0:
			s += (t * 60)
		case 1:
			s += t
		}
	}
	return s, errors.New("problem")
}
