package cah

import (
	"Nie-Mand/cah-poc/internal/core/base"
	"Nie-Mand/cah-poc/internal/core/deck"
	"Nie-Mand/cah-poc/internal/core/player"
	"fmt"
)

func (c *CardsAgainstHumanity) NextRound() {
	c.round++
	c.players[c.round%len(c.players)].SetCzar()
}

func (c *CardsAgainstHumanity) PickWinner(_player base.Player, card base.Card) error {
	__player := _player.(*player.DirtyPlayer)
	if !__player.IsCzar() {
		return fmt.Errorf("%s is not the czar", __player.Name())
	}

	_card, err := card.(*deck.ShittyCard)
	if !err {
		return fmt.Errorf("Invalid card type: %v", err)
	}

	winner := _card.Owner()
	winner.IncrementScore(1)

	return nil
}

func (c *CardsAgainstHumanity) CleanRound() error {
	c.reviewQueue.Clear()

	for _, player := range c.players {
		if !player.IsCzar() {
			c.Deal(player)
		} else {
			player.UnsetCzar()
		}
	}

	return nil
}
