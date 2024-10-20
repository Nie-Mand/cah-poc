package deck

import "math/rand"

type Deck []*ShittyCard


func NewDeck() (Deck, Deck, error) {
	blackCards := []*ShittyCard{
		NewBlackCard("Why can't I sleep at night?"),
		NewBlackCard("I got 99 problems but _ ain't one."),
		NewBlackCard("Monalisa's smile, _"),
	}

	whiteCards := []*ShittyCard{
		NewWhiteCard("A big black..."),
		NewWhiteCard("It hurts"),
		NewWhiteCard("69"),
		NewWhiteCard("It's a trap!"),
		NewWhiteCard("I'm a banana!"),
		NewWhiteCard("What is love?"),
		NewWhiteCard("Wow!"),
		NewWhiteCard("OG Kush"),
		NewWhiteCard("Really, Mom?"),
		NewWhiteCard("I'm Batman"),
		NewWhiteCard("Rainbow Dash"),
		NewWhiteCard("I'm a unicorn!"),
		NewWhiteCard("I'm a wizard!"),
		NewWhiteCard("I'm a ninja!"),
		NewWhiteCard("I'm a pirate!"),
		NewWhiteCard("I'm a robot!"),
		NewWhiteCard("I'm a vampire!"),
		NewWhiteCard("I'm a werewolf!"),
		NewWhiteCard("I'm a zombie!"),
		NewWhiteCard("I'm a ghost!"),
		NewWhiteCard("I'm a mermaid!"),
		NewWhiteCard("I'm a fairy!"),
		NewWhiteCard("I'm a princess!"),
		NewWhiteCard("I'm a prince!"),
		NewWhiteCard("I'm a king!"),
		NewWhiteCard("I'm a queen!"),
		NewWhiteCard("I'm a joker!"),
		NewWhiteCard("I'm a jester!"),
		NewWhiteCard("I'm a knight!"),
		NewWhiteCard("I'm a dragon!"),
		NewWhiteCard("I'm a monster!"),
		NewWhiteCard("I'm a hero!"),
		NewWhiteCard("I'm a villain!"),
		NewWhiteCard("I'm a god!"),
		NewWhiteCard("I'm a goddess!"),
		NewWhiteCard("I'm a demon!"),
		NewWhiteCard("I'm a devil!"),
		NewWhiteCard("I'm a saint!"),
		NewWhiteCard("I'm a sinner!"),
		NewWhiteCard("I'm a lover!"),
		NewWhiteCard("I'm a fighter!"),
		NewWhiteCard("I'm a winner!"),
		NewWhiteCard("I'm a loser!"),
		NewWhiteCard("I'm a leader!"),
		NewWhiteCard("I'm a follower!"),
		NewWhiteCard("I'm a teacher!"),
		NewWhiteCard("I'm a student!"),
		NewWhiteCard("I'm a parent!"),
		NewWhiteCard("I'm a child!"),
	}

	return blackCards, whiteCards, nil
}

func (d *Deck) Head() *ShittyCard {
	return (*d)[0]
}

func (d *Deck) Draw() *ShittyCard {
	card := (*d)[0]
	*d = (*d)[1:]
	return card
}

func (d *Deck) Shuffle() {
	for i := range *d {
		j := rand.Intn(i + 1)
		(*d)[i], (*d)[j] = (*d)[j], (*d)[i]
	}
}
