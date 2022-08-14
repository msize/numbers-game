package human_vs_human

import (
	"crypto/rand"
	"fmt"

	"github.com/numbers-game/consts"
	"github.com/numbers-game/permutations"
	"github.com/numbers-game/score"
	"github.com/numbers-game/types"
)

func Run() {
	fmt.Println("Human vs Human mode")
	numbers := permutations.Generate()
	fmt.Print("Your hidden number: ")
	hidden := fetchRundomNumber(numbers)
	fmt.Println(hidden)
	youFirst := youFirst()
	var yourScores types.Scores
	var yourScore, opponentScore types.Score
	for i := 2; !score.Win(yourScore) || !score.Win(opponentScore); i++ {
		if i%2 == 0 {
			fmt.Printf("========== %d ==========\n", i/2)
			fmt.Println(!score.Win(yourScore) || !score.Win(opponentScore))
		}
		if i%2 == 0 && youFirst || i%2 != 0 && !youFirst {
			yourScores = yourTurn(numbers, yourScores)
			numbers = score.Sift(numbers, yourScores)
			yourScore = yourScores[len(yourScores)-1]
		} else {
			opponentScore = opponentTurn(hidden)
		}
	}
	fmt.Println("=======================")
}

func yourTurn(numbers types.Numbers, yourScores types.Scores) types.Scores {
	leftNumbers := len(numbers)
	if leftNumbers == 0 {
		fmt.Println("Your turn: No suggestions")
		return yourScores
	}
	fmt.Print("Your turn: ")
	calledNumber := fetchRundomNumber(numbers)
	fmt.Print(calledNumber)
	fmt.Printf(" (%d)\n", leftNumbers)
	fmt.Print("Your score: ")
	if leftNumbers == 1 {
		fmt.Println(" 4 4")
		return score.AppendScore(yourScores, types.Score{Number: calledNumber, Guessed: consts.Digits, Postions: consts.Digits})
	}
	return score.AppendScore(yourScores, scanScore(calledNumber))
}

func opponentTurn(hidden types.Number) types.Score {
	fmt.Print("Opponent turn: ")
	opponent := scanNumber()
	opponentScore := score.Calc(opponent, hidden)
	fmt.Printf("Opponent score: %d-%d\n", opponentScore.Guessed, opponentScore.Postions)
	return opponentScore
}

func fetchRundomNumber(numbers types.Numbers) types.Number {
	return numbers[randInt(len(numbers))]
}

func randInt(len int) uint {
	RandomCrypto, _ := rand.Prime(rand.Reader, 128)
	return uint(RandomCrypto.Int64()) % uint(len)
}

func scanNumber() types.Number {
	var result types.Number
	for i := 0; i < consts.Digits; i++ {
		fmt.Scanf("%d", &result[i])
	}
	return result
}

func scanScore(calledNumber types.Number) types.Score {
	var result types.Score
	fmt.Scanf("%d", &result.Guessed)
	fmt.Scanf("%d", &result.Postions)
	result.Number = calledNumber
	return result
}

func youFirst() bool {
	fmt.Print("Who is first? [0-you]: ")
	var input int
	fmt.Scanf("%d", &input)
	return input == 0
}
