package main

import (
	"bytes"
	"reflect"
	"testing"
	"time"
)

func TestCountdown(t *testing.T) {
	t.Run("sleep btw 2 prints", func(t *testing.T) {
		spySleepPrint := &SpyCountdownOperations{}
		Countdown(spySleepPrint, spySleepPrint)
		want := []string{
			write,
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
		}
		if !reflect.DeepEqual(spySleepPrint.Calls, want) {
			t.Errorf("got calls in order %v but wanted calls %v", spySleepPrint, want)
		}
	})
	t.Run("prints 321Go!", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		spySleeper := &SpyCountdownOperations{}

		Countdown(buffer, spySleeper)

		got := buffer.String()
		want :=
			`3
2
1
Go!`
		if got != want {
			t.Errorf(" Got : %q , wanted : %q", got, want)
		}
	})

}

func TestConfigurableSleeper(t *testing.T) {
	sleepTime := 5 * time.Second
	spyTime := &SpyTime{}
	sleeper := &ConfigurableSleeper{sleepTime, spyTime.Sleep}
	sleeper.Sleep()
	if sleepTime != spyTime.durationSlept {
		t.Errorf("should have slept for %v but slept for %v", sleepTime, spyTime.durationSlept)
	}
}
