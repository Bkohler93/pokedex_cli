package repl

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/bkohler93/pokedexcli/internal/api/pokeapi"
	"github.com/bkohler93/pokedexcli/internal/app/pokedex_cli/commands"
	"github.com/bkohler93/pokedexcli/internal/repository"
)

func StartRepl(pokeRepo *repository.PokeRepository) {
	s := bufio.NewScanner(os.Stdin)
	pokedex := map[string]pokeapi.RespPokemonData{}

	for {
		fmt.Print("pokedex > ")
		s.Scan()

		input := s.Text()
		input = cleanInput(input)
		if len(input) == 0 {
			continue
		}

		args := strings.Fields(input)
		command := args[0]

		c, ok := commands.GetCommands()[command]
		if !ok {
			fmt.Println("Invalid command. Type 'help' to see available commands.")
			continue
		}

		var err error
		if len(args) == 1 {
			err = c.Callback(pokeRepo, pokedex)
		} else {
			err = c.Callback(pokeRepo, pokedex, args[1:]...)
		}
		if err != nil {
			fmt.Println("Error...", err)
		}
	}
}

func cleanInput(in string) string {
	in = strings.ToLower(in)
	in = strings.Trim(in, " ")
	return in
}
