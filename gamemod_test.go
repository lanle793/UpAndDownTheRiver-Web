package main

import "testing"

func TestGetPointsForRoundCorrectGuess(t *testing.T) {
	guess := 5
	wins := 5
	expected := 15
	actual := getPointsForRound(guess, wins)
	if actual != expected {
		t.Errorf("Result was incorrect, got: %d, want: %d.", actual, expected)
	}
}

func TestGetPointsForRoundCorrect(t *testing.T) {
	guess := 5
	wins := 3
	expected := 0
	actual := getPointsForRound(guess, wins)
	if actual != expected {
		t.Errorf("Result was incorrect, got: %d, want: %d.", actual, expected)
	}
}