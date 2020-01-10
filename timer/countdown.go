package timer

import (
	"fmt"
	"time"
)

//CountDown accepts an integer to count down from in seconds. The function will print the remaining seconds left, every second, and then prints "Go!"
func CountDown(n int) error {
	var err error
	for n > 0 {
		_, err = fmt.Println(n)
		if (err != nil) {
			return err
		}
		time.Sleep(time.Second)
		n--
	}
	_, err = fmt.Println("Go!")
	return err
}