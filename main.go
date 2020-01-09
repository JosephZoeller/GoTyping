package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func init() {
}

func main() {
	preface()
	countDown(3)
	runTest()
}

func preface() {
	fmt.Println("Welcome to my typing speed test. This program will count down from 3, and then you will type words as fast as you can. " +
		"When you're ready, press the return key to begin.")
	read()
}

func runTest() {
	start := time.Now()

	inputString, readError := read()

	elapsed := time.Since(start)

	if readError != nil {
		fmt.Println(readError)
		os.Exit(-1)
	}

	fmt.Println("Seconds to complete:", elapsed.Seconds())

	userWords := strings.Split(inputString, " ")
	var wordCount float64 = 0
	for i := 0; i < len(userWords); i++ {
		if userWords[i] != "" {
			wordCount++
		}
	}
	fmt.Println("Words per second:", wordCount/elapsed.Seconds())
	fmt.Println("Words per minute:", wordCount/elapsed.Seconds() * 60)	
	fmt.Print("\n")
}

func read() (string, error) {
	reader := bufio.NewReader(os.Stdin)
	str, err := reader.ReadString('\n')
	return str, err
}

func countDown(i int) {
	for i > 0 {
		fmt.Println(i)
		time.Sleep(time.Second)
		i--
	}
	println("Go!")
}
