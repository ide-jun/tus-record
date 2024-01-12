package clock

import "time"

type Clock interface {
	Now() time.Time
}

type RealClock struct{}

var clk = &RealClock{}

func (s *RealClock) Now() time.Time {
	return time.Now()
}

func New() Clock {
	return clk
}
