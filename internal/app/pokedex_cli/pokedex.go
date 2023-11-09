package pokedexcli

import (
	"github.com/bkohler93/pokedexcli/internal/app/pokedex_cli/constants"
	"github.com/bkohler93/pokedexcli/internal/app/pokedex_cli/repl"
	"github.com/bkohler93/pokedexcli/internal/repository"
)

func RunApp() {
	repo := repository.NewPokeRepository(constants.ApiTimeout, constants.CacheReapTimeout)

	repl.StartRepl(repo)
}
