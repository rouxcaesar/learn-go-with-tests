package main

import (
  "io"
  "fmt"
  "os"
  "time"
)

func main() {
  sleeper := &ConfigurableSleeper{1 * time.Second, time.Sleep}
  Countdown(os.Stdout, sleeper)
}

const finalWord = "Go!"
const countdownStart = 3

type Sleeper interface {
  Sleep()
}

func Countdown(out io.Writer, sleeper Sleeper) {
  for i := countdownStart; i > 0; i-- {
    sleeper.Sleep()
    fmt.Fprintln(out, i)
  }

  sleeper.Sleep()
  fmt.Fprintln(out, finalWord)
}
