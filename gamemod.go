package main

import "fmt"

func Init() {
	fmt.Println("Let the game begin!")
	fmt.Println(getPointsForRound(1, 1))
}
 
func getPointsForRound(guess int, wins int) int {
	if wins == guess {
		return 10 + wins
	} else {
		return 0
	}
}