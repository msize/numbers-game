package human_vs_machine

import (
	"fmt"

	"github.com/numbers-game/consts"
	"github.com/numbers-game/game/common"
	"github.com/numbers-game/permutations"
	"github.com/numbers-game/score"
	"github.com/numbers-game/types"
)

func Run() {
	fmt.Println("Human vs Machine mode")
	numbers := permutations.Generate()
	hidden := common.FetchRundomNumber(numbers)
	youFirst := common.YouFirst()
	var machineScores types.Scores
	var machineScore, humanScore types.Score
	for i := 2; !score.Win(humanScore) || !score.Win(machineScore); i++ {
		if i%2 == 0 {
			fmt.Printf("========== %d ==========\n", i/2)
		}
		if i%2 == 0 && youFirst || i%2 != 0 && !youFirst {
			humanScore = humanTurn(hidden)
		} else {
			machineScores = machineTurn(numbers, machineScores)
			numbers = score.Sift(numbers, machineScores)
			machineScore = machineScores[len(machineScores)-1]
		}
	}
	fmt.Println("=======================")
}

func humanTurn(hidden types.Number) types.Score {
	fmt.Print("Human turn: ")
	human := common.ScanNumber()
	humanScore := score.Calc(human, hidden)
	fmt.Printf("Human score: %d-%d\n", humanScore.Guessed, humanScore.Postions)
	return humanScore
}

func machineTurn(numbers types.Numbers, machineScores types.Scores) types.Scores {
	leftNumbers := len(numbers)
	if leftNumbers == 0 {
		fmt.Println("Machine turn: No suggestions")
		return machineScores
	}
	fmt.Print("Machine turn: ")
	calledNumber := common.FetchRundomNumber(numbers)
	fmt.Print(calledNumber)
	fmt.Printf(" (%d)\n", leftNumbers)
	fmt.Print("Machine score: ")
	if leftNumbers == 1 {
		fmt.Println(" 4 4")
		return score.AppendScore(machineScores, types.Score{Number: calledNumber, Guessed: consts.Digits, Postions: consts.Digits})
	}
	return score.AppendScore(machineScores, common.ScanScore(calledNumber))
}
