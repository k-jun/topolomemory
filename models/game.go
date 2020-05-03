package models

import (
	"sync"
	"errors"
)

type IGame interface {
	Start() error
	Draw() error
	Hit(a int, b int) (bool, error)
	Status() []Card
}

func NewGame() IGame {
	return &game{}
}

type game struct {
	deck  []Card
	field []Card
	mutex sync.Mutex
}

func (g *game) Start() error {
	g.deck = deck
	g.field = field
	return nil
}

func (g *game) Draw() error {
	g.mutex.Lock()
	defer g.mutex.Unlock()
	if len(g.deck) == 0 {
		return errors.New("no card on deck")
	}
	card := g.deck[0]
	g.deck = g.deck[1:len(g.deck)]
	g.field = append(g.field, card)
	return nil

}

func (g *game) Hit(a int, b int) (bool, error) {
	g.mutex.Lock()
	defer g.mutex.Unlock()
	aIdx, err := g.find(a)
	if err != nil {
		return false, err
	}

	bIdx, err := g.find(b)
	if err != nil {
		return false, err
	}

	acard := g.field[aIdx]
	bcard := g.field[bIdx]
	flag := acard.TopologyHoles == bcard.TopologyHoles && acard.TopologyParts == bcard.TopologyParts

	if flag {
		g.field = append(g.field[:aIdx], g.field[aIdx+1:]...)
		g.field = append(g.field[:bIdx], g.field[bIdx+1:]...)
	}

	return flag, nil
}

func (g *game) Status() []Card {
	return g.field
}

func (g *game) find(id int) (int, error) {
	for idx, c := range g.field {
		if c.Id == id {
			return idx, nil
		}

	}
	return -1, errors.New("not found")
}


