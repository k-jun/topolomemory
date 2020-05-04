package models

import (
	"errors"
	"math/rand"
	"sync"
	"time"
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
	g.mutex.Lock()
	defer g.mutex.Unlock()
	g.deck = deck
	g.field = field
	g.shuffleDeck()
	// discard one card
	g.deck = g.deck[1:len(g.deck)]
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
	acard, err := g.find(a)
	if err != nil {
		return false, err
	}

	bcard, err := g.find(b)
	if err != nil {
		return false, err
	}

	flag := acard.TopologyHoles == bcard.TopologyHoles && acard.TopologyParts == bcard.TopologyParts

	if flag {
		g.removeTwo(acard, bcard)
	}

	return flag, nil
}

func (g *game) Status() []Card {
	g.mutex.Lock()
	// g.shuffleField()
	defer g.mutex.Unlock()
	return g.field
}

func (g *game) shuffleDeck() error {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(g.deck), func(i, j int) { g.deck[i], g.deck[j] = g.deck[j], g.deck[i] })
	return nil
}

func (g *game) shuffleField() error {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(g.field), func(i, j int) { g.field[i], g.field[j] = g.field[j], g.field[i] })
	return nil
}

func (g *game) removeTwo(a *Card, b *Card) error {
	nf := []Card{}
	for _, v := range g.field {
		if v.Id == a.Id || v.Id == b.Id {
			continue
		}
		nf = append(nf, v)
	}
	g.field = nf
	return nil
}

func (g *game) find(id int) (*Card, error) {
	for _, c := range g.field {
		if c.Id == id {
			return &c, nil
		}

	}
	return nil, errors.New("not found")
}
