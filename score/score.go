package score

import (
	"github.com/numbers-game/consts"
	"github.com/numbers-game/types"
)

type Number types.Number

type Score types.Score

const digits = consts.Digits

func Win(score Score) bool {
	return score[0] == consts.Digits && score[1] == consts.Digits
}

func Equals(first Score, second Score) bool {
	return first[0] == second[0] && first[1] == second[1]
}

func Calc(called Number, hidden Number) Score {
	return Score{guessed(called, hidden), postions(called, hidden)}
}

func guessed(called Number, hidden Number) int8 {
	result := int8(0)
	for i := 0; i < digits; i++ {
		for j := 0; j < digits; j++ {
			if called[i] == hidden[j] {
				result++
				break
			}
		}
	}
	return result
}

func postions(called Number, hidden Number) int8 {
	result := int8(0)
	for i := 0; i < digits; i++ {
		if called[i] == hidden[i] {
			result++
		}
	}
	return result
}
