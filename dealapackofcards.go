package piscine

import "fmt"

func DealAPackOfCards(deck []int) {
	// We know there are 12 cards and 4 players
	cardsPerPlayer := 3
	playerCount := 4
	index := 0

	for i := 1; i <= playerCount; i++ {
		fmt.Printf("Player %d: ", i)
		for j := 0; j < cardsPerPlayer; j++ {
			fmt.Printf("%d", deck[index])
			index++
			if j < cardsPerPlayer-1 {
				fmt.Printf(", ")
			}
		}
		fmt.Printf("\n")
	}
}
