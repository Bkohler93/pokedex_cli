package commands

import (
	"fmt"

	"github.com/bkohler93/pokedexcli/internal/pokedex"
	"github.com/bkohler93/pokedexcli/internal/repository"
)

func helpCallback(pokeRepo *repository.PokeRepository, pokedex pokedex.Pokedex, args ...string) error {
	fmt.Println("Welcome to the Pokedex!\n\nUsage:")
	for _, c := range GetCommands() {
		fmt.Printf("'%s' - %s\n", c.name, c.description)
	}

	return nil
}
