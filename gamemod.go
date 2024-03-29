package main

import (
	"fmt"
	"sort"
)

func Init() {
	fmt.Println("Let the game begin!")
	fmt.Println(getPointsForRound(1, 1))
}
 
func getPointsForRound(guess int, wins int) int {
	var points int

	if wins == guess {
		if guess == 0 {
			points = 5
		} else {
			points = wins + 10
		}
	} else {
		points = wins - guess
		if points < 0 {
			points = 0
		}
	}

	return points
}

func getNumCardsOfRound(lastRoundNumCards int, isIncrement bool) int {
	if isIncrement {
		return lastRoundNumCards + 1
	} else {
		return lastRoundNumCards - 1
	}
}

func getTrickWinner(playerList []Player, trumpSuit string, leadSuit string) Player {
	sort.Slice(playerList, func(i, j int) bool {
		card1 := playerList[i].CurrentCard
		card2 := playerList[j].CurrentCard

		if card1.Suit == card2.Suit {
			return card1.Value > card2.Value
		} else if card1.Suit == trumpSuit || card2.Suit == trumpSuit {
			return card1.Suit == trumpSuit
		} else if card1.Suit == leadSuit || card2.Suit == leadSuit {
			return card1.Suit == leadSuit
		}

		return false
	})
	return playerList[0]
}

func getGameWinner(playerList []Player) Player {
	sort.Slice(playerList, func(i, j int) bool {
		return playerList[i].Points > playerList[j].Points
	})
	return playerList[0]
}