package main

import "fmt"

func main() {
	fmt.Println("Welcome to the card game: Up and Down the River!!!")
	fmt.Println(getPointsForRound(1, 1))
}

func getPointsForRound(guess int, wins int) int {
	if wins == guess {
		return 10 + wins
	} else {
		return 0
	}
}