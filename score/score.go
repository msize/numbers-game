package score

import (
	"github.com/numbers-game/consts"
	"github.com/numbers-game/types"
)

func Win(score types.Score) bool {
	return score.Guessed == consts.Digits && score.Positions == consts.Digits
}

func AppendScore(scores types.Scores, score types.Score) types.Scores {
	return append(scores, score)
}

func Calc(called types.Number, hidden types.Number) types.Score {
	return types.Score{
		Number:    called,
		Guessed:   guessed(called, hidden),
		Positions: positions(called, hidden),
	}
}

func Sift(numbers types.Numbers, scores types.Scores) types.Numbers {
	var result types.Numbers
	for _, num := range numbers {
		if fitNumber(num, scores) {
			result = append(result, num)
		}
	}
	return result
}

func fitNumber(number types.Number, scores types.Scores) bool {
	for _, called := range scores {
		if !equals(called, Calc(called.Number, number)) {
			return false
		}
	}
	return true
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

func positions(called types.Number, hidden types.Number) int8 {
	result := int8(0)
	for i := 0; i < consts.Digits; i++ {
		if called[i] == hidden[i] {
			result++
		}
	}
	return result
}

func equals(first types.Score, second types.Score) bool {
	return first.Guessed == second.Guessed && first.Positions == second.Positions
}
