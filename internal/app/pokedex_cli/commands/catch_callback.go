package commands

import (
	"errors"
	"fmt"

	catchcalculator "github.com/bkohler93/pokedexcli/internal/app/pokedex_cli/catch_calculator"
	"github.com/bkohler93/pokedexcli/internal/pokedex"
	"github.com/bkohler93/pokedexcli/internal/repository"
)

func catchCallback(repo *repository.PokeRepository, pokedex pokedex.Pokedex, args ...string) error {
	if len(args) < 1 {
		return errors.New("enter a Pokemon you want to catch")
	}

	pokemonName := args[0]

	pokemonData, ok, err := repo.GetPokemonByName(pokemonName)
	if err != nil {
		return err
	}

	if !ok {
		return fmt.Errorf("no Pokemon found with the name %s", pokemonName)
	}

	fmt.Printf("Throwing a Pokeball at %s... ", pokemonName)

	if isCaught := catchcalculator.EvaluateCatchOutcome(pokemonData.BaseExperience); !isCaught {
		fmt.Printf("%s escaped!\n", pokemonData.Name)
		return nil
	}

	if _, ok := pokedex[pokemonName]; !ok {
		pokedex[pokemonName] = pokemonData
	}

	fmt.Printf("%s was caught!\n", pokemonName)

	return nil
}
