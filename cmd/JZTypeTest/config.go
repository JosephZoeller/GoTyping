// Package main project-0 Typing Test for 200106-uta-go
package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	tbutil "github.com/JosephZoeller/project-0/pkg/termboxutil"
	"github.com/JosephZoeller/project-0/pkg/timer"
	"github.com/nsf/termbox-go"
	"log"
	"os"
	"runtime"
	"strconv"
	"strings"
)

const configfile = "config.json"
const LOGFILE string = "log.txt"

var LogFile *os.File

var user *string
var freestyle, debug, cheat *bool
var duration int

var sentences []string

type configJson struct {
	SentenceFile     string `json:"sentencePrompts"`
	DefaultDuration  string `json:"duration"`
	DefaultFreestyle bool   `json:"freestyle"`
	DefaultVerbose   bool   `json:"verbose"`
	DefaultCheat     bool   `json:"cheat"`
}

func config() {
	initLog()

	initTermbox()

	log.Println("[config]: Reading " + configfile)
	c := configJson{}
	cFile, confErr := os.Open(configfile)
	defer cFile.Close()

	if confErr != nil {
		log.Println("[config]: Error: " + confErr.Error())
		log.Println("[config]: Resorting to backup defaults.")
		_, y := tbutil.Write(0, 0, termbox.ColorRed, tbutil.COLDEF, "Error: "+configfile+" could not be opened. Using backup defaults.")
		tbutil.Write(0, y+2, tbutil.COLDEF, tbutil.COLDEF, "Press any key to continue...")
		setFallBackDefaults(&c)
		tbutil.KeyContinue(false)
	} else {
		confErr = json.NewDecoder(cFile).Decode(&c)
		if confErr != nil {
			log.Println("[config]: Error: " + confErr.Error())
			log.Println("[config]: Resorting to backup defaults.")
			_, y := tbutil.Write(0, 0, termbox.ColorRed, tbutil.COLDEF, "Error: "+configfile+" could not be decoded. Using backup defaults.")
			tbutil.Write(0, y+2, tbutil.COLDEF, tbutil.COLDEF, "Press any key to continue...")
			setFallBackDefaults(&c)
			tbutil.KeyContinue(false)
		}
	}

	log.Println("[config]: Parsing arguments")
	user = flag.String("u", strings.Title(getDefaultName()), "User - Defaults to the Operating System's current username.")
	freestyle = flag.Bool("f", c.DefaultFreestyle, "Freestyle - Removes the writing prompt. The user can type without restriction and accuracy won't be measured.")
	debug = flag.Bool("v", c.DefaultVerbose, "Verbose - Displays under-the-hood details during the test.")
	cheat = flag.Bool("c", c.DefaultCheat, "Cheat - Fudges the test results to impress your peers.")
	durStr := flag.String("d", c.DefaultDuration, "Duration - The length of time that the typing test will last. Format as <Minutes>:<Seconds>")
	flag.Parse()

	buildSentences(c.SentenceFile)

	arbitrateDuration(c.DefaultDuration, *durStr)

	log.Println("[config]: Configuration complete")
}

func setFallBackDefaults(c *configJson) {
	c.SentenceFile = "sentences.txt"
	c.DefaultDuration = "0:30"
	c.DefaultFreestyle = false
	c.DefaultVerbose = false
	c.DefaultCheat = false
}

func initLog() {
	var confErr error
	LogFile, confErr = os.Create("log.txt")
	if confErr != nil {
		log.Fatal(confErr)
	}
	log.SetOutput(LogFile)
}

func buildSentences(filename string) { // call after parsing args
	if *freestyle == true {
		return
	}

	log.Println("[config]: Building writing prompts")
	f, senfilerr := os.Open(filename)
	if senfilerr != nil {
		*freestyle = true
		log.Println("[config]: Error: " + senfilerr.Error())
		log.Println("[config]: Defaulting to freestyle mode.")
		_, y := tbutil.Write(0, 0, termbox.ColorRed, tbutil.COLDEF, "Error: "+filename+" could not be opened. Defaulting to freestyle mode.")
		tbutil.Write(0, y+2, tbutil.COLDEF, tbutil.COLDEF, "Press any key to continue...")
		tbutil.KeyContinue(false)
	} else {
		sc := bufio.NewScanner(f)
		for sc.Scan() {
			sentences = append(sentences, sc.Text())
		}
		if len(sentences) < 2 {
			sentences = append(sentences,
				fmt.Sprintf("There are not enough prompts in the %s file.", filename),
				"This is a default prompt. The prompt file was found lacking.")
		}
	}
}

func initTermbox() {
	log.Println("[config]: Initializing termbox-go...")
	err := termbox.Init()
	if err != nil {
		log.Fatalln(err)
	}
	termbox.SetInputMode(termbox.InputEsc)
}

func getDefaultName() string {
	log.Println("[config]: Getting default username")
	var userEnvVar string
	if runtime.GOOS == "windows" {
		userEnvVar = os.Getenv("USERNAME")
	} else if runtime.GOOS == "linux" {
		userEnvVar = os.Getenv("USER")
	} else {
		userEnvVar = "User"
	}
	return userEnvVar
}

func arbitrateDuration(jsonDefault, argStr string) {
	log.Println("[config]: Parsing test duration")
	var arbErr error
	// set default
	duration, arbErr = timer.ParseCountDown(jsonDefault)
	if arbErr != nil {
		//default backup
		duration = 30
		log.Println("[config]: Error: " + arbErr.Error())
		log.Printf("[config]: Cannot parse %s's duration string \"%s\" (desired format: <# minutes>:<# seconds>).", configfile, jsonDefault)
	}

	durArg, arbErr := timer.ParseCountDown(argStr)
	if arbErr != nil {
		// use default
		log.Println("[config]: Error: " + arbErr.Error())
		log.Printf("[config]: Cannot parse the duration argument \"%s\" (desired format: <# minutes>:<# seconds>). Defaulting to %ds.", argStr, duration)
		_, y := tbutil.Write(0, 0, termbox.ColorRed, tbutil.COLDEF, "Error: Cannot parse duration argument \""+argStr+"\" (desired format: <# minutes>:<# seconds>). Defaulting to "+strconv.Itoa(duration)+"s.")
		tbutil.Write(0, y+2, tbutil.COLDEF, tbutil.COLDEF, "Press any key to continue...")
		tbutil.KeyContinue(false)
	} else {
		// use argument duration
		duration = durArg
		log.Printf("[config]: Duration set to %ds.", duration)
	}
}
