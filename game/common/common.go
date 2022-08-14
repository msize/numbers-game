package common

import (
	"crypto/rand"
	"fmt"

	"github.com/numbers-game/consts"
	"github.com/numbers-game/types"
)

func YouFirst() bool {
	fmt.Print("Who is first? [0-you]: ")
	var input int
	fmt.Scanf("%d", &input)
	return input == 0
}

func FetchRundomNumber(numbers types.Numbers) types.Number {
	return numbers[randInt(len(numbers))]
}

func randInt(len int) uint {
	RandomCrypto, _ := rand.Prime(rand.Reader, 128)
	return uint(RandomCrypto.Int64()) % uint(len)
}

func ScanScore(calledNumber types.Number) types.Score {
	var result types.Score
	fmt.Scanf("%d", &result.Guessed)
	fmt.Scanf("%d", &result.Postions)
	result.Number = calledNumber
	return result
}

func ScanNumber() types.Number {
	var result types.Number
	for i := 0; i < consts.Digits; i++ {
		fmt.Scanf("%d", &result[i])
	}
	return result
}
