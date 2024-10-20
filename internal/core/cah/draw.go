package cah

import (
	"Nie-Mand/cah-poc/internal/core/base"
	"Nie-Mand/cah-poc/internal/core/deck"
	"Nie-Mand/cah-poc/internal/core/player"
	"fmt"
)

func (c *CardsAgainstHumanity) Draw() (base.Card, error) {
	black := c.blackDeck.Draw()
	return black, nil
}

func (c *CardsAgainstHumanity) Deal(_player base.Player) error {
	dirtyPlayer := _player.(*player.DirtyPlayer)
	dirtyPlayer.Draw(c.whiteDeck.Draw())
	return nil
}

func (c *CardsAgainstHumanity) DealN(player base.Player, count int) error {
	for i := 0; i < count; i++ {
		c.Deal(player)
	}

	return nil
}

func (c *CardsAgainstHumanity) Choose(_player base.Player, card base.Card) error {
	__player := _player.(*player.DirtyPlayer)
	_card := card.(*deck.ShittyCard)

	if !__player.HasCard(_card) {
		return fmt.Errorf("%s does not have the card %s", __player.Name(), _card.Content())
	}

	c.reviewQueue.Push(_card)

	err := __player.RemoveCard(_card)
	if err != nil {
		return err
	}

	_card.SetOwner(__player)

	return nil
}
