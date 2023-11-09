package commands

import (
	"github.com/bkohler93/pokedexcli/internal/pokedex"
	"github.com/bkohler93/pokedexcli/internal/repository"
)

type command struct {
	name        string
	description string
	Callback    func(repo *repository.PokeRepository, pokedex pokedex.Pokedex, args ...string) error
}

func GetCommands() map[string]command {
	return map[string]command{
		"help": {
			name:        "help",
			description: "Displays information on how to use the Pokedex",
			Callback:    helpCallback,
		},
		"exit": {
			name:        "exit",
			description: "Exits the Pokedex",
			Callback:    exitCallback,
		},
		"map": {
			name:        "map",
			description: "Displays location areas in the Pokemon world (20 each time)",
			Callback:    mapCallback,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the previous 20 locations in the Pokemon world",
			Callback:    mapBCallback,
		},
		"explore": {
			name:        "explore",
			description: "Use 'explore [valid area-name]' to view information on the chosen area",
			Callback:    exploreCallback,
		},
		"catch": {
			name:        "catch",
			description: "Use 'explore [valid pokemon name]' to try to catch the chosen pokemon",
			Callback:    catchCallback,
		},
		"inspect": {
			name:        "inspect",
			description: "'inspect [valid pokemon name] displays details about the pokemon",
			Callback:    inspectCallback,
		},
		"pokedex": {
			name:        "pokedex",
			description: "lists all caught pokemon by name",
			Callback:    pokedexCallback,
		},
	}
}
