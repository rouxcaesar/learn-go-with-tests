package main

import "fmt"

const spanish = "Spanish"
const spanishHelloPrefix = "Hola, "
const englishHelloPrefix = "Hello, "

func main() {
  fmt.Println(Hello("Chris", ""))
}

func Hello(name, language string) string {
  if name == "" {
    name = "World"
  }

  if language == "Spanish" {
    return spanishHelloPrefix + name
  }

  return englishHelloPrefix + name
}
