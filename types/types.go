package types

import "github.com/numbers-game/consts"

type Number [consts.Digits]int8

type Numbers []Number

type Score struct {
	Number    Number
	Guessed   int8
	Positions int8
}

type Scores []Score
