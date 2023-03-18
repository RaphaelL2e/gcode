package utils

import (
	"fmt"
	"time"
)

/**
This is a tool for counting time.
*/

type Timer struct {
	check     bool
	startTime time.Time
}

func NewTimer() Timer {
	return Timer{check: true}
}

func (t *Timer) StaticsTime() {
	if t.check {
		t.startTime = time.Now()
		t.check = false
	} else {
		fmt.Println(time.Now().Sub(t.startTime))
		t.startTime = time.Now()
	}
}
