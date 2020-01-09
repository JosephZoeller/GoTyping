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
	fmt.Println("Welcome to my typing speed test. This program will count down from 3, and then it will measure how fast you can type words. " +
		"When you're ready, press the return key to begin and end.")
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

	printStats(elapsed, countWords(inputString))
}

func printStats(elapsed time.Duration, wordCount int) {
	wordCountFl := float64(wordCount)
	fmt.Println("Seconds to complete:", elapsed.Seconds())
	fmt.Println("Words per second:", wordCountFl/elapsed.Seconds())
	fmt.Println("Words per minute:", wordCountFl/elapsed.Seconds()*60)
}

func countWords(words string) int {
	wordArray := strings.Split(words, " ")
	wordCount := 0
	for i := 0; i < len(wordArray); i++ {
		if wordArray[i] != "" {
			wordCount++
		}
	}
	return wordCount
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
