package commands

import (
	"errors"
	"fmt"

	"github.com/bkohler93/pokedexcli/internal/pokedex"
	"github.com/bkohler93/pokedexcli/internal/repository"
)

func mapCallback(pokeRepo *repository.PokeRepository, pokedex pokedex.Pokedex, args ...string) error {
	locations, err := pokeRepo.GetNextLocations()
	if err != nil {
		return errors.New("error retrieving locations")
	}

	for _, l := range locations.Results {
		fmt.Println(l.Name)
	}

	return nil
}

func mapBCallback(pokeRepo *repository.PokeRepository, pokedex pokedex.Pokedex, args ...string) error {
	locations, err := pokeRepo.GetPrevLocations()
	if err != nil {
		return err
	}

	for _, l := range locations.Results {
		fmt.Println(l.Name)
	}

	return nil
}
