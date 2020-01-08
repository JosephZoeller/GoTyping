package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func cleanHouse() {
	fmt.Print("\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n")
}

func main() {
	//cleanHouse()
	countDown()

	start := time.Now()

	reader := bufio.NewReader(os.Stdin)

	inputString, readError := reader.ReadString('\n')

	elapsed := time.Since(start)

	if readError == nil {
		fmt.Println("You wrote: " + inputString) // includes the newline
		fmt.Printf("It took %s to enter this string.", elapsed)
	} else {
		fmt.Println(readError)
	}

	fmt.Print("\n")
}

func countDown() {
	println("3")
	time.Sleep(time.Second)
	println("2")
	time.Sleep(time.Second)
	println("1")
	time.Sleep(time.Second)
	println("Go!")
}
