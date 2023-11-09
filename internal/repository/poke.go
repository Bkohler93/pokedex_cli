package repository

import (
	"errors"
	"time"

	"github.com/bkohler93/pokedexcli/internal/api/pokeapi"
)

type PokeRepository struct {
	client           *pokeapi.Client
	NextLocationsURL *string
	PrevLocationsURL *string
}

func NewPokeRepository(requestTimeout time.Duration, cacheReapTimeout time.Duration) *PokeRepository {
	return &PokeRepository{
		client:           pokeapi.NewPokeApiClient(requestTimeout, cacheReapTimeout),
		NextLocationsURL: nil,
		PrevLocationsURL: nil,
	}
}

func (r *PokeRepository) GetNextLocations() (pokeapi.RespLocationsList, error) {
	var locations pokeapi.RespLocationsList
	var err error

	locations, err = r.client.GetLocations(r.NextLocationsURL)
	if err != nil {
		return locations, err
	}

	r.NextLocationsURL = locations.Next
	r.PrevLocationsURL = locations.Previous

	return locations, nil
}

func (r *PokeRepository) GetPrevLocations() (pokeapi.RespLocationsList, error) {
	var locations pokeapi.RespLocationsList

	if r.PrevLocationsURL == nil {
		return locations, errors.New("you are on the first page of locations")
	}

	locations, err := r.client.GetLocations(r.PrevLocationsURL)
	if err != nil {
		return locations, err
	}

	r.NextLocationsURL = locations.Next
	r.PrevLocationsURL = locations.Previous

	return locations, nil
}

func (r *PokeRepository) GetLocationData(name string) (pokeapi.RespLocationData, error) {
	var locationData pokeapi.RespLocationData

	locationData, err := r.client.GetLocation(name)
	if err != nil {
		return locationData, err
	}

	return locationData, nil
}

func (r *PokeRepository) GetPokemonByName(name string) (pokeapi.RespPokemonData, bool, error) {
	var pokemon pokeapi.RespPokemonData
	var isValidName bool
	var err error

	pokemon, isValidName, err = r.client.GetPokemon(name)
	if err != nil {
		return pokemon, isValidName, err
	}
	if !isValidName {
		return pokemon, isValidName, err
	}

	return pokemon, isValidName, err
}
