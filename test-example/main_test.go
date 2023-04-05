package main

import (
	"testing"

	"fyne.io/fyne/test"
)

func TestGreeting(t *testing.T) {
	out, in := makeUI()
	test.Type(in, "Andy")

	if out.Text != "Hello Andy!" {
		t.Error(t.Name(), ": ", "Incorrect initial greeting")
	}
}
