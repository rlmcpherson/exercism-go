package clock

import "fmt"

const TestVersion = 2

type Clock struct {
	min int
}

func Time(hour, minute int) Clock {
	c := Clock{(60*hour + minute) % 1440}
	if c.min < 0 {
		c.min += 1440
	}
	return c
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.min/60, c.min%60)
}

func (c Clock) Add(min int) Clock {
	c.min = (c.min + min) % 1440
	if c.min < 0 {
		c.min += 1440
	}
	return c
}
