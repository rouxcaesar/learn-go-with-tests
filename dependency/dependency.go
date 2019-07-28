package main

import (
  "fmt"
  "net/http"
  "io"
)

func main() {
  http.ListenAndServe(":5000", http.HandlerFunc(MyGreeterHandler))
}

func Greet(writer io.Writer, name string) {
  fmt.Fprintf(writer, "Hello, %s", name)
}

func MyGreeterHandler(w http.ResponseWriter, r *http.Request) {
  Greet(w, "world")
}
