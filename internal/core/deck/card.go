package deck

import (
	"Nie-Mand/cah-poc/internal/core/base"
	"Nie-Mand/cah-poc/internal/core/player"
)

type CardType string

const (
	BlackCard CardType = "black"
	WhiteCard CardType = "white"
)

type ShittyCard struct {
	kind               CardType
	content            string
	howManyCardsToDraw int
	owner              *player.DirtyPlayer
}

func (c *ShittyCard) Content() string {
	return c.content
}

func (c ShittyCard) Owner() base.Player {
	return c.owner
}

func (c *ShittyCard) SetOwner(p base.Player) {
	c.owner = p.(*player.DirtyPlayer)
}

func NewBlackCard(text string) *ShittyCard {
	return &ShittyCard{
		kind:               BlackCard,
		content:            text,
		howManyCardsToDraw: 1,
	}
}

func NewBlackCardWithDraw(text string, draw int) *ShittyCard {
	return &ShittyCard{
		kind:               BlackCard,
		content:            text,
		howManyCardsToDraw: draw,
	}
}

func NewWhiteCard(text string) *ShittyCard {
	return &ShittyCard{
		kind:               WhiteCard,
		content:            text,
		howManyCardsToDraw: 0,
	}
}
