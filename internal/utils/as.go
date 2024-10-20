package utils

import "Nie-Mand/cah-poc/internal/core/base"

func As[X base.CardGame](game base.CardGame) X {
	return game.(X)
}
