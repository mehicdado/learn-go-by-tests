package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

//Sleeper guarantee existence of Sleep method implementation.
type Sleeper interface {
	Sleep()
}

//DefaultSleeper sleep for 1 sec always.
type DefaultSleeper struct{}

//Sleep implementation of DefaultSleeper sleep always for 1 second.
func (d *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

//ConfigurableSleeper give possibility to configure seep time.
type ConfigurableSleeper struct {
	duration time.Duration
	sleep    func(time.Duration)
}

//Sleep for given duration.
func (c *ConfigurableSleeper) Sleep() {
	c.sleep(c.duration)
}

const finalWord = "Go!"
const countdownStart = 3

//Countdown print 3,2,1,Go! to given writer and sleep between each write.
func Countdown(writer io.Writer, sleeper Sleeper) {
	for i := countdownStart; i > 0; i-- {
		sleeper.Sleep()
		fmt.Fprintln(writer, i)
	}
	sleeper.Sleep()
	fmt.Fprint(writer, finalWord)
}

func main() {
	sleeper := &ConfigurableSleeper{2 * time.Second, time.Sleep}
	Countdown(os.Stdout, sleeper)
}
