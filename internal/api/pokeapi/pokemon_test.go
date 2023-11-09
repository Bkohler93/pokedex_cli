package pokeapi

import (
	"testing"
	"time"
)

func TestGetPokemon(t *testing.T) {
	pokemonName := "pikachu"
	client := NewPokeApiClient(time.Second*10, time.Second*10)

	pokemon, isValid, err := client.GetPokemon("pikachu")
	if err != nil {
		t.Errorf("%s expected no error", pokemonName)
	}

	if !isValid {
		t.Errorf("%s should be a valid pokemon", pokemonName)
	}

	if pokemon.Name != pokemonName {
		t.Errorf("%s is not %s", pokemon.Name, pokemonName)
	}
}
