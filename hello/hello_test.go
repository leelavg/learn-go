package main

import "testing"

func TestHello(t *testing.T){

    assertCorrectMessage := func(t *testing.T, got, want string) {
        t.Helper()      // Gives exact line of failure
        if got != want{
            t.Errorf("Got %q want %q", got, want)
        }
    }

    t.Run("saying hello to people", func(t *testing.T){
        got := Hello("Leela", "")
        want := "Hello, Leela"
        assertCorrectMessage(t, got, want)
    })

    t.Run("say 'Hello, World' when and empty string is supplied", func(t *testing.T){
        got := Hello("","")
        want := "Hello, World"
        assertCorrectMessage(t, got, want)
    })

    t.Run("in Spanish", func(t *testing.T){
        got := Hello("Leela", "Spanish")
        want := "Hola, Leela"
        assertCorrectMessage(t, got, want)
    })

    t.Run("in French", func(t *testing.T){
        got := Hello("Leela", "French")
        want := "Bonjour, Leela"
        assertCorrectMessage(t, got, want)
    })
}

