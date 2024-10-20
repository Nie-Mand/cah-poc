package main

import (
	"Nie-Mand/cah-poc/internal/core/base"
	"Nie-Mand/cah-poc/internal/core/cah"
	"Nie-Mand/cah-poc/internal/core/deck"
	"Nie-Mand/cah-poc/internal/core/player"
	"Nie-Mand/cah-poc/internal/utils"
	"fmt"
	"math/rand/v2"
)

func handleError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	players := []*player.DirtyPlayer{
		player.NewDirtyPlayer("Alice"),
		player.NewDirtyPlayer("Bob"),
		player.NewDirtyPlayer("Charlie"),
	}

	blackDeck, whiteDeck, err := deck.NewDeck()
	handleError(err)

	blackDeck.Shuffle()
	whiteDeck.Shuffle()

	var game base.CardGame = cah.NewCardsAgainstHumanity(players, whiteDeck, blackDeck)

	// Initial deck
	for _, player := range players {
		err := game.DealN(player, cah.DRAW_CARDS_COUNT)
		handleError(err)
	}

	roundsCount := 3
	for i := 0; i < roundsCount; i++ {
		game.NextRound()
		fmt.Println("Round:", i)

		black, err := game.Draw()
		fmt.Println("Black card:", black.Content())

		czar := utils.As[*cah.CardsAgainstHumanity](game).GetJudge()
		fmt.Println("Czar:", czar.Name())

		for _, player := range players {
			if player.IsCzar() {
				continue
			}

			randomChoice := rand.IntN(len(player.Deck()))
			card := player.Deck()[randomChoice]
			err := game.Choose(player, card)
			handleError(err)
			fmt.Println("Player:", player.Name(), "chose:", card.Content())
		}

		// Next round
		q, err := utils.As[*cah.CardsAgainstHumanity](game).GetQueue(*czar)
		handleError(err)

		randomChoice := rand.IntN(q.Size())
		fmt.Println("Winner choice:", randomChoice)

		var card base.Card
		for i := 0; i <= randomChoice; i++ {
			card, err = q.Pop()
			handleError(err)
		}

		fmt.Println("Winner card:", card.Content())
		err = game.PickWinner(czar, card)
		handleError(err)

		for _, player := range players {
			fmt.Println(player.Name(), "score:", player.Score())
		}

		err = game.CleanRound()
		handleError(err)
	}

	// Winner
	winners := game.Winner()
	if len(winners) > 1 {
		fmt.Println("Draw")
	} else {
		fmt.Println("Winner:", winners[0].Name())
	}
}
