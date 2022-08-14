package score

import (
	"github.com/numbers-game/consts"
	"github.com/numbers-game/types"
)

func Win(score types.Score) bool {
	return score[0] == consts.Digits && score[1] == consts.Digits
}

func Equals(first types.Score, second types.Score) bool {
	return first[0] == second[0] && first[1] == second[1]
}

func Calc(called types.Number, hidden types.Number) types.Score {
	return types.Score{guessed(called, hidden), postions(called, hidden)}
}

func guessed(called types.Number, hidden types.Number) int8 {
	result := int8(0)
	for i := 0; i < consts.Digits; i++ {
		for j := 0; j < consts.Digits; j++ {
			if called[i] == hidden[j] {
				result++
				break
			}
		}
	}
	return result
}

func postions(called types.Number, hidden types.Number) int8 {
	result := int8(0)
	for i := 0; i < consts.Digits; i++ {
		if called[i] == hidden[i] {
			result++
		}
	}
	return result
}
