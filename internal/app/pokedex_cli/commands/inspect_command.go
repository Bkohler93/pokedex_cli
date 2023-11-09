package commands

import (
	"errors"
	"fmt"

	"github.com/bkohler93/pokedexcli/internal/pokedex"
	"github.com/bkohler93/pokedexcli/internal/repository"
)

func inspectCallback(repo *repository.PokeRepository, pokedex pokedex.Pokedex, args ...string) error {
	if len(args) < 1 {
		return errors.New("'inspect' expects a pokemon name as its second argument")
	}

	pokemonName := args[0]

	if pokemon, ok := pokedex[pokemonName]; !ok {
		return fmt.Errorf("you have yet to catch a %s", pokemonName)
	} else {
		fmt.Printf("%s Details:\n", pokemonName)
		fmt.Printf("height: %d\n", pokemon.Height)
		fmt.Printf("weight: %d\n", pokemon.Weight)

		fmt.Println()
		fmt.Println("Stats")
		for _, stat := range pokemon.Stats {
			fmt.Printf("%s: base %d, effort %d\n", stat.Stat.Name, stat.BaseStat, stat.Effort)
		}
		fmt.Println()

		fmt.Print("Types: ")
		for _, t := range pokemon.Types {
			fmt.Printf("%s", t.Type.Name)
		}
		fmt.Println()
	}

	return nil
}
