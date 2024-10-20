package player

import (
	"Nie-Mand/cah-poc/internal/core/base"
	"fmt"
)

type DirtyPlayer struct {
	name   string
	score  int
	deck   []base.Card
	isCzar bool
}

func NewDirtyPlayer(name string) *DirtyPlayer {
	return &DirtyPlayer{
		name:   name,
		score:  0,
		deck:   make([]base.Card, 0),
		isCzar: false,
	}
}

func (p *DirtyPlayer) Draw(cards ...base.Card) {
	p.deck = append(p.deck, cards...)
}

func (p *DirtyPlayer) IncrementScore(count int) {
	p.score += count
}

func (p DirtyPlayer) HasCard(card base.Card) bool {
	for _, c := range p.deck {
		if c == card {
			return true
		}
	}
	return false
}

func (p DirtyPlayer) RemoveCard(card base.Card) error {
	for i, c := range p.deck {
		if c == card {
			p.deck = append(p.deck[:i], p.deck[i+1:]...)
			return nil
		}
	}

	return fmt.Errorf("%s does not have the card %s", p.Name(), card.Content())
}

func (p DirtyPlayer) Score() int {
	return p.score
}

func (p DirtyPlayer) Name() string {
	return p.name
}

func (p DirtyPlayer) Deck() []base.Card {
	return p.deck
}

func (p *DirtyPlayer) SetCzar() {
	p.isCzar = true
}

func (p *DirtyPlayer) UnsetCzar() {
	p.isCzar = false
}

func (p DirtyPlayer) IsCzar() bool {
	return p.isCzar
}
