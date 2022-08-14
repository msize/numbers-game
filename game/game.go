package game

import (
	"crypto/rand"
	"fmt"

	"github.com/numbers-game/consts"
	"github.com/numbers-game/permutations"
	"github.com/numbers-game/score"
	"github.com/numbers-game/types"
)

func Run() {
	numbers := permutations.Generate()
	hidden := fetchRundomNumber(numbers)
	if showMachineHidden() {
		fmt.Print("Machine hidden number: ")
		fmt.Println(hidden)
	}
	humanFirst := humanFirst()
	var machineScores types.Scores
	var machineScore, humanScore types.Score
	for i := 2; !score.Win(humanScore) || !score.Win(machineScore); i++ {
		if i%2 == 0 {
			fmt.Printf("========== %d ==========\n", i/2)
		}
		if i%2 == 0 && humanFirst || i%2 != 0 && !humanFirst {
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
	human := scanNumber()
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
	calledNumber := fetchRundomNumber(numbers)
	fmt.Print(calledNumber)
	fmt.Printf(" (%d)\n", leftNumbers)
	fmt.Print("Machine score: ")
	if leftNumbers == 1 {
		fmt.Println(" 4 4")
		return score.AppendScore(machineScores, types.Score{Number: calledNumber, Guessed: consts.Digits, Postions: consts.Digits})
	}
	return score.AppendScore(machineScores, scanScore(calledNumber))
}

func humanFirst() bool {
	fmt.Print("Who is first? [0 - human]: ")
	var input int
	fmt.Scanf("%d", &input)
	return input == 0
}

func showMachineHidden() bool {
	fmt.Print("Would you like see machine's hidden? [1 - yes]: ")
	var input int
	fmt.Scanf("%d", &input)
	return input == 1
}

func fetchRundomNumber(numbers types.Numbers) types.Number {
	return numbers[randInt(len(numbers))]
}

func randInt(len int) uint {
	RandomCrypto, _ := rand.Prime(rand.Reader, 128)
	return uint(RandomCrypto.Int64()) % uint(len)
}

func scanScore(calledNumber types.Number) types.Score {
	var result types.Score
	fmt.Scanf("%d", &result.Guessed)
	fmt.Scanf("%d", &result.Postions)
	result.Number = calledNumber
	return result
}

func scanNumber() types.Number {
	var result types.Number
	for i := 0; i < consts.Digits; i++ {
		fmt.Scanf("%d", &result[i])
	}
	return result
}
