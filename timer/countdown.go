package timer

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

//CountDown accepts an integer to count down from in seconds. The function will print the remaining seconds left, every second.
func CountDown(n int) error {
	var err error
	for n > 0 {
		_, err = fmt.Println(n)
		if err != nil {
			return err
		}
		time.Sleep(time.Second)
		n--
	}
	return nil
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