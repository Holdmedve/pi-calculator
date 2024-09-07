package main

import (
	"testing"
	"strconv"
)

func TestCalculatePi(t *testing.T) {
	got := strconv.FormatFloat(CalculatePi(), 'g', -1, 64)[:6]
	want := "3.1415"

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

