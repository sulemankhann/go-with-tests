package main

import (
	"bytes"
	"testing"
)

func TestGree(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Suleman")

	got := buffer.String()
	want := "Hello, Suleman"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
