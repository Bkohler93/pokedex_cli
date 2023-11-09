package commands

import (
	"errors"
	"fmt"

	"github.com/bkohler93/pokedexcli/internal/pokedex"
	"github.com/bkohler93/pokedexcli/internal/repository"
)

func pokedexCallback(repo *repository.PokeRepository, pokedex pokedex.Pokedex, args ...string) error {

	if len(pokedex) == 0 {
		return errors.New("you have yet to catch any pokemon. Try using 'catch' to catch one first")
	}

	for _, v := range pokedex {
		fmt.Println(v.Name)
	}

	return nil
}
