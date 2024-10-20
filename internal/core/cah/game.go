package cah

import (
	"Nie-Mand/cah-poc/internal/core/base"
	"Nie-Mand/cah-poc/internal/core/deck"
	"Nie-Mand/cah-poc/internal/core/player"
	"fmt"
)

type CardsAgainstHumanity struct {
	players     player.DirtyPlayers
	blackDeck   deck.Deck
	whiteDeck   deck.Deck
	reviewQueue *deck.ReviewQueue
	round       int
}

const (
	DRAW_CARDS_COUNT = 10
)

func NewCardsAgainstHumanity(players []*player.DirtyPlayer, whiteDeck deck.Deck, blackDeck deck.Deck) *CardsAgainstHumanity {
	return &CardsAgainstHumanity{
		players:     players,
		blackDeck:   blackDeck,
		whiteDeck:   whiteDeck,
		reviewQueue: deck.NewReviewQueue(len(players)),
		round:       0,
	}
}

func (c *CardsAgainstHumanity) GetQueue(_player player.DirtyPlayer) (*deck.ReviewQueue, error) {
	if _player.IsCzar() {
		return c.reviewQueue, nil
	}

	return nil, fmt.Errorf("%s is not the czar", _player.Name())
}

func (c *CardsAgainstHumanity) GetJudge() *player.DirtyPlayer {
	for _, player := range c.players {
		if player.IsCzar() {
			return player
		}
	}

	return nil
}

func (c *CardsAgainstHumanity) Winner() []base.Player {
	max := 0
	for _, player := range c.players {
		if player.Score() > max {
			max = player.Score()
		}
	}

	winners := make([]base.Player, 0, len(c.players))
	for _, player := range c.players {
		if player.Score() == max {
			winners = append(winners, player)
		}
	}

	return winners
}
