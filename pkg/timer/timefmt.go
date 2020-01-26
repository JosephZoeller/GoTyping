// Package timer project-0 Typing Test for 200106-uta-go
package timer

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

// ParseCountDown attempts to parse a time format into seconds.
// accepts a string in the format of <minutes>:<seconds>,
// and returns an int which represents that time in seconds.
func ParseCountDown(a string) (int, error) {
	s := 0
	aDlm := strings.Split(a, ":")
	for i, a := range aDlm {
		t, e := strconv.Atoi(a)
		if e != nil {
			return -1, e
		}
		switch i {
		case 0:
			s += (t * 60)
		case 1:
			s += t
		}
	}
	return s, nil
}

func GetDateFmt() string {
	y, m, d := time.Now().Date()
	return fmt.Sprintf("%d/%d/%d", m, d, y)
}
