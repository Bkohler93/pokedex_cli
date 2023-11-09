package commands

import (
	"fmt"
	"os"

	"github.com/bkohler93/pokedexcli/internal/pokedex"
	"github.com/bkohler93/pokedexcli/internal/repository"
)

func exitCallback(pokeRepo *repository.PokeRepository, pokedex pokedex.Pokedex, args ...string) error {
	fmt.Println("Have a nice day!")
	os.Exit(1)
	return nil
}
