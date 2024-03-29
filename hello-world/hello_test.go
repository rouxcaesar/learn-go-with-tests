package main

import "testing"

func TestHello(t *testing.T) {
  assertCorrectMessage := func(t *testing.T, got, want string) {
    t.Helper()
    if got != want {
      t.Errorf("got '%s' want '%s'", got, want)
    }
  }

  t.Run("saying hello to people", func(t *testing.T) {
    got := Hello("Chris", "")
    want := "Hello, Chris"
    assertCorrectMessage(t, got, want)
  })

  t.Run("Say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
    got := Hello("", "")
    want := "Hello, World"
    assertCorrectMessage(t, got, want)
  })

  t.Run("in Spanish", func(t *testing.T) {
    got := Hello("Elodie", "Spanish")
    want := "Hola, Elodie"
    assertCorrectMessage(t, got, want)
  })

  t.Run("in French", func(t *testing.T) {
    got := Hello("Frances", "French")
    want := "Bonjour, Frances"
    assertCorrectMessage(t, got, want)
  })

  t.Run("in Afrikaans", func(t *testing.T) {
    got := Hello("Megan", "Afrikaans")
    want := "Hallo, Megan"
    assertCorrectMessage(t, got, want)
  })
}
