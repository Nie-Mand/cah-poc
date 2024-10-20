package base

type Player interface {
	Score() int
	IncrementScore(count int)
	Name() string
	Deck() []Card
	Draw(cards ...Card)
	HasCard(card Card) bool
	RemoveCard(card Card) error
}
