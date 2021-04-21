package main

import (
	"regexp"
	"testing"
)

func TestWelcome(t *testing.T) {
	want := regexp.MustCompile(`\b` + "TicTacToe" + `\b`)
	msg, err := Welcome()
	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Welcome() = %q, %v, want match for %#q, nil`, msg, err, want)
	}
}
