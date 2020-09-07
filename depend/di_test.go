package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Leela")

	got := buffer.String()
	want := "Hello, Leela"

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
