package timer

import (
	"fmt"
	"time"
)

//CountDown accepts an integer to count down from in seconds. The function will print the remaining seconds left, every second, and then prints "Go!"
func CountDown(n int) {
	for n > 0 {
		fmt.Println(n)
		time.Sleep(time.Second)
		n--
	}
	println("Go!")
}
