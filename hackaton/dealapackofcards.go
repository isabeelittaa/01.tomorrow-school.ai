package piscine

import "fmt"

func DealAPackOfCards(deck []int) {
	numPlayers := 4
	cardsPerPlayer := 3
	for i := 0; i < numPlayers; i++ {
		fmt.Printf("Player %d: ", i+1)
		start := i * cardsPerPlayer
		end := start + cardsPerPlayer
		for j := start; j < end; j++ {
			fmt.Printf("%d", deck[j])
			if j < end-1 {
				fmt.Printf(", ")
			}
		}
		fmt.Printf("\n")
	}
}
