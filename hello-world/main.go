package main

import "fmt"

const spanish = "Spanish"
const french = "French"
const afrikaans = "Afrikaans"
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "
const afrikaansHelloPrefix = "Hallo, "
const englishHelloPrefix = "Hello, "

func main() {
  fmt.Println(Hello("Chris", ""))
}

func Hello(name, language string) string {
  if name == "" {
    name = "World"
  }

  return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
  switch language {
  case french:
    prefix = frenchHelloPrefix
  case spanish:
    prefix = spanishHelloPrefix
  case afrikaans:
    prefix = afrikaansHelloPrefix
  default:
    prefix = englishHelloPrefix
  }

  return
}
