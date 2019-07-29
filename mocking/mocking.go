package main

import (
  "io"
  "fmt"
  "os"
  "time"
)

func main() {
  Countdown(os.Stdout)
}

const finalWord = "Go!"
const countdownStart = 3

type Sleeper interface {
  Sleep()
}

func Countdown(out io.Writer) {
  for i := countdownStart; i > 0; i-- {
    time.Sleep(1 * time.Second)
    fmt.Fprintln(out, i)
  }

  time.Sleep(1 * time.Second)
  fmt.Fprintln(out, finalWord)
}
