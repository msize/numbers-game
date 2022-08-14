package permutations

import (
	"github.com/numbers-game/consts"
	"github.com/numbers-game/types"
)

func Generate() types.Numbers {
	result := types.Numbers{}
	currentNumber := firstNumber()
	for hasNext := true; hasNext; currentNumber, hasNext = next(currentNumber) {
		result = append(result, currentNumber)
	}
	return result
}

func firstNumber() types.Number {
	var result types.Number
	var i int8
	for i = 0; i < consts.Digits; i++ {
		result[i] = i
	}
	return result
}

func next(number types.Number) (types.Number, bool) {
	return inc(number, consts.Digits-1)
}

func inc(number types.Number, position int) (types.Number, bool) {
	if position < 0 {
		return number, false
	}
	number[position] = nextFree(number, position)
	if number[position] != -1 {
		return number, true
	}
	result, hasNext := inc(number, position-1)
	result[position] = nextFree(result, position)
	return result, hasNext
}

func nextFree(number types.Number, position int) int8 {
	result := number[position] + 1
	for ; result < 10; result++ {
		if !occupied(number, result, position) {
			return result
		}
	}
	return -1
}

func occupied(number types.Number, digit int8, position int) bool {
	for i := 0; i < consts.Digits; i++ {
		if i != position && number[i] == digit {
			return true
		}
	}
	return false
}
