package main

import (
	"baccarat-bot/simulation"
)

func main() {
	simulation.Game(simulation.Bet{Amount: 1000, Hand: simulation.Player})
}
