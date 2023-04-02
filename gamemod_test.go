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

func TestGetPointsForRoundCorrectGuessZero(t *testing.T) {
	guess := 0
	wins := 0
	expected := 5
	actual := getPointsForRound(guess, wins)
	if actual != expected {
		t.Errorf("Result was incorrect, got: %d, want: %d.", actual, expected)
	}
}

func TestGetPointsForRoundFewerWinsThanGuess(t *testing.T) {
	guess := 5
	wins := 3
	expected := 0
	actual := getPointsForRound(guess, wins)
	if actual != expected {
		t.Errorf("Result was incorrect, got: %d, want: %d.", actual, expected)
	}
}

func TestGetPointsForRoundMoreWinsThanGuess(t *testing.T) {
	guess := 5
	wins := 7
	expected := 2
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

func TestGetTrickWinner(t *testing.T) {
	player1 := Player{"Jane", 50, Card{"Heart", 12}}
	player2 := Player{"John", 34, Card{"Club", 4}}
	player3 := Player{"April", 77, Card{"Spade", 7}}
	playerList := []Player{player1, player2, player3}

	trickWinner := getTrickWinner(playerList, "Spade", "Heart")

	if trickWinner != player3 {
		t.Errorf("Incorrect winner, got: %+v, want: %+v", trickWinner, player3)
	}
}

func TestGetGameWinner(t *testing.T) {
	player1 := Player{"Jane", 50, Card{"Heart", 12}}
	player2 := Player{"John", 34, Card{"Club", 4}}
	player3 := Player{"April", 77, Card{"Spade", 7}}
	playerList := []Player{player1, player2, player3}
	
	winner := getGameWinner(playerList)

	if winner != player3 {
		t.Errorf("Incorrect winner, got: %+v, want: %+v", winner, player3)
	}
}