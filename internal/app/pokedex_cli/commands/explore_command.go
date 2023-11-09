package commands

import (
	"errors"
	"fmt"

	"github.com/bkohler93/pokedexcli/internal/pokedex"
	"github.com/bkohler93/pokedexcli/internal/repository"
)

func exploreCallback(pokeRepo *repository.PokeRepository, pokedex pokedex.Pokedex, args ...string) error {

	if len(args) == 0 {
		return errors.New("requires second argument... example 'explore oreburgh-mine-1f'")
	}

	location := args[0]

	locationData, err := pokeRepo.GetLocationData(location)
	if err != nil {
		return errors.New("unable to retrieve location data")
	}

	for _, v := range locationData.PokemonEncounters {
		fmt.Println(v.Pokemon.Name)
	}

	return nil
}
