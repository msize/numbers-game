package human_vs_human

import (
	"fmt"

	"github.com/numbers-game/consts"
	"github.com/numbers-game/game/common"
	"github.com/numbers-game/permutations"
	"github.com/numbers-game/score"
	"github.com/numbers-game/types"
)

func Run() {
	fmt.Println("Human vs Human mode")
	numbers := permutations.Generate()
	fmt.Print("Your hidden number: ")
	hidden := common.FetchRundomNumber(numbers)
	fmt.Println(hidden)
	youFirst := common.YouFirst()
	var yourScores types.Scores
	var yourScore, opponentScore types.Score
	for i := 2; !score.Win(yourScore) || !score.Win(opponentScore); i++ {
		if i%2 == 0 {
			fmt.Printf("========== %d ==========\n", i/2)
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
	calledNumber := common.FetchRundomNumber(numbers)
	fmt.Print(calledNumber)
	fmt.Printf(" (%d)\n", leftNumbers)
	fmt.Print("Your score: ")
	if leftNumbers == 1 {
		fmt.Println(" 4 4")
		return score.AppendScore(yourScores, types.Score{Number: calledNumber, Guessed: consts.Digits, Postions: consts.Digits})
	}
	return score.AppendScore(yourScores, common.ScanScore(calledNumber))
}

func opponentTurn(hidden types.Number) types.Score {
	fmt.Print("Opponent turn: ")
	opponent := common.ScanNumber()
	opponentScore := score.Calc(opponent, hidden)
	fmt.Printf("Opponent score: %d-%d\n", opponentScore.Guessed, opponentScore.Postions)
	return opponentScore
}
