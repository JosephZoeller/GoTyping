# GoTyping - CLI Typing Test
This interactive Go program allows the user to test their typing skills using [termbox-go](https://github.com/nsf/termbox-go) for real-time input and statistics.

## Getting started
Upon running the application, the user is immediately met with a preface message with instructions for how to take the test. By default, the test will prompt the user with sentences to type, for 30 seconds. Once completed, the program will display statistics about the user's performance.

This typing evaluation is customizable. The user can modify the test duration, the writing prompts, the testing mode and other miscellanious variables. The default settings (such as the 30 second duration) can be customized through the config.json file, or passed in directly as arguments.

### Installation
Install this go package with `go get -u github.com/JosephZoeller/project-0` and termbox-go with `go get -u github.com/nsf/termbox-go`

## Command-line args / config.json
The typing test accepts numerous arguments:

**-d ##:##**

  * Duration - The length of time that the typing test will last. Format as \<Minutes>:\<Seconds> (default "0:30")

**-f**

  * Freestyle - Removes the writing prompt. The user can type without restriction and accuracy won't be measured.

**-u <name>**

  * User - Defaults to the Operating System's current username. (default "Joseph")

**-v**

  * Verbose - Displays under-the-hood details during the test.

**-c**   

  * Cheat - Fudges the test results to impress your peers.

## Features Roadmap
- [x] Documentation
- [x] Unit testing (TODO simulate user input)
- [x] Benchmarking
- [x] Logging
- [X] API Library (Stopwatch, countdown)
- [x] CLI flags (Timer, freestyle/sentence reference)
- [x] Environment variables (user)
- [x] Concurrency (Timer)
- [x] Data Persistance (reference sentences, TODO user averages/statistics)
- [x] HTTP API (TODO user averages/statistics)

## Presentation
- [x] 5 minute live demonstration
- [x] Slides & visual aides

## Found a bug?

Please submit a bug report to GitHub with as much detail as possible. Please include the log.txt if applicable.
