package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func init() {
}

func main() {
	countDown(3)
	beginTest()
}

func beginTest() {
	start := time.Now()

	inputString, readError := read()

	elapsed := time.Since(start)

	if readError == nil {
		fmt.Println("You wrote: " + inputString) // includes the newline
		fmt.Printf("It took %s to enter this string.", elapsed)
	} else {
		fmt.Println(readError)
	}

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
