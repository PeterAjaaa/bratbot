package simulation

import (
	"fmt"
	"math/rand"
)

// Simulates a loop of the game
func Game(bot_bet Bet) {
	var values = []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K"}
	var suites = []string{"♤", "♡", "♢", "♧"}
	var decks []string

	decks = createDecks(values, suites, decks)
	shuffle(decks)

	fmt.Println("Decks: ", decks)
	fmt.Println("Decks length: ", len(decks))

}

// Fisher-Yates shuffle, specifically for 8 decks baccarat
func shuffle(decks []string) {
	for i := len(decks) - 1; i > 0; i-- {
		var random = rand.Intn(i + 1)
		decks[random], decks[i] = decks[i], decks[random]
	}
}

// Simulate the use of 8 decks
func createDecks(values []string, suites []string, decks []string) []string {
	for i := 0; i < 8; i++ {
		for _, value := range values {
			for _, suite := range suites {
				decks = append(decks, fmt.Sprintf("%v%v", value, suite))
			}
		}
	}
	return decks
}
