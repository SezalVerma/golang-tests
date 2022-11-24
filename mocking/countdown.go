package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

type Sleeper interface {
	Sleep()
}

// to make sleep time configurable
type ConfigurableSleeper struct {
	duration time.Duration
	// func(time.Duration ) same as time.Sleep
	sleep func(time.Duration)
}

func (c *ConfigurableSleeper) Sleep() {
	c.sleep(c.duration)
}

// implement configurableSleeper
type SpyTime struct {
	durationSlept time.Duration
}

func (spy *SpyTime) Sleep(duration time.Duration) {
	spy.durationSlept = duration
}

// spysleeper - mock that mocks the behaviour of sleep
// it just counts rather than calling time.Sleep() .
// less time to exec in test
type SpyCountdownOperations struct {
	Calls []string
}

const (
	write = "write"
	sleep = "sleep"
)

func (s *SpyCountdownOperations) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

func (s *SpyCountdownOperations) Write(w []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

// our sleeper
type DefaultSleeper struct{}

func (d *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

// countdown
func Countdown(w io.Writer, s Sleeper) {
	for i := 3; i >= 1; i-- {
		fmt.Fprintln(w, i)
		s.Sleep()
	}
	fmt.Fprint(w, "Go!")
}

// main
func main() {
	sleeper := &DefaultSleeper{}
	Countdown(os.Stdout, sleeper)

	sleeper2 := &ConfigurableSleeper{1 * time.Second, time.Sleep}
	Countdown(os.Stdout, sleeper2)
}
