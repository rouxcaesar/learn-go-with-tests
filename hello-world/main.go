package main

import "fmt"

const spanish = "Spanish"
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "
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
  
  if language == "French" {
    return frenchHelloPrefix + name
  }

  return englishHelloPrefix + name
}
