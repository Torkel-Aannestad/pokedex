package main

import (
	"errors"
	"sync"

	"github.com/torkelaannestad/pokedex/internal/pokeapi"
)

type Pokedex struct {
	Pokedex map[string]pokeapi.Pokemon
	mu      *sync.Mutex
}

func NewPokedex() Pokedex {
	return Pokedex{
		Pokedex: map[string]pokeapi.Pokemon{},
	}
}

func (p Pokedex) Add(name string, pokemon pokeapi.Pokemon) {
	p.mu.Lock()
	defer p.mu.Unlock()

	_, exists := p.Pokedex[name]
	if !exists {
		p.Pokedex[name] = pokemon
	}
}

func (p Pokedex) Get(name string) (pokeapi.Pokemon, error) {
	p.mu.Lock()
	defer p.mu.Unlock()

	v, exists := p.Pokedex[name]
	if exists {
		return v, nil
	} else {
		return pokeapi.Pokemon{}, errors.New("not in Pokedex")
	}
}
func (p Pokedex) GetAll() ([]pokeapi.Pokemon, error) {
	p.mu.Lock()
	defer p.mu.Unlock()

	pokedex := []pokeapi.Pokemon{}
	for _, v := range p.Pokedex {
		pokedex = append(pokedex, v)
	}
	return pokedex, nil
}

func (p Pokedex) Delete(name string) {
	p.mu.Lock()
	defer p.mu.Unlock()

	delete(p.Pokedex, name)
}
