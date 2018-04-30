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

const finalWord = "Go!"
const countdownStart = 3

func Countdown(out io.Writer, sleeper Sleeper) {
	for i := countdownStart; i > 0; i-- {
		time.Sleep(1 * time.Second)
		fmt.Fprintln(out, i)
	}

	fmt.Fprint(out, finalWord)
}

type ConfigurableSleeper struct {
	Duration time.Duration
}

func (o *ConfigurableSleeper) Sleep() {
	time.Sleep(o.Duration)
}

func main() {
	sleeper := &ConfigurableSleeper{Duration: 1 * time.Second}
	Countdown(os.Stdout, sleeper)
}
