package base

type Card interface {
	Content() string
	Owner() Player
}
