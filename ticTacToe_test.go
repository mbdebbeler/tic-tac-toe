package main

import (
	"bytes"
	"testing"
)

func TestAskHuman(t *testing.T) {
	testInput := "X"
	var stdin bytes.Buffer
	stdin.Write([]byte(testInput + "\n"))

	result, err := askHuman("chooseMarker", &stdin)
	if result != testInput || err != nil {
		t.Fatalf(`askHuman() = %v, want match for %q, nil`, result, testInput)
	}
}
