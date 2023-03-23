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

func TestGetPointsForRoundIncorrectGuess(t *testing.T) {
	guess := 5
	wins := 3
	expected := 0
	actual := getPointsForRound(guess, wins)
	if actual != expected {
		t.Errorf("Result was incorrect, got: %d, want: %d.", actual, expected)
	}
}

func TestGetNumCardsOfRound(t *testing.T) {
	numCards := getNumCardsOfRound(10, false)
	if numCards != 9 {
		t.Errorf("Number of cards was incorrect, got: %d, want: 9", numCards)
	}

	numCards = getNumCardsOfRound(6, true)
	if numCards != 7 {
		t.Errorf("Number of cards was incorrect, got: %d, want: 7", numCards)
	}
}