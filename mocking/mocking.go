package main

import (
  "io"
  "fmt"
  "os"
  "time"
)

func main() {
  sleeper := &DefaultSleeper{}
  Countdown(os.Stdout, sleeper)
}

const finalWord = "Go!"
const countdownStart = 3

type DefaultSleeper struct {}

type Sleeper interface {
  Sleep()
}

func (d *DefaultSleeper) Sleep() {
  time.Sleep(1 * time.Second)
}

func Countdown(out io.Writer, sleeper Sleeper) {
  for i := countdownStart; i > 0; i-- {
    sleeper.Sleep()
    fmt.Fprintln(out, i)
  }

  sleeper.Sleep()
  fmt.Fprintln(out, finalWord)
}
