package base

type CardGame interface {
	Draw() (Card, error)
	Choose(player Player, card Card) error
	Deal(player Player) error
	DealN(player Player, count int) error
	NextRound()
	PickWinner(player Player, card Card) error
	CleanRound() error
	Winner() []Player
}
