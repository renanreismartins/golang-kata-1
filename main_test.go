package main

import "testing"

func TestAbs(t *testing.T) {
	got := welcomeMessage()
	if got != "Hi" {
		t.Errorf("welcomeMessage() = %s; want Hi", got)
	}
}