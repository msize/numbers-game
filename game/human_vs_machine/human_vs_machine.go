package human_vs_machine

import (
	"fmt"

	"github.com/numbers-game/permutations"
)

func Run() {
	fmt.Println("Human vs Machine mode")
	numbers := permutations.Generate()
	for _, number := range numbers {
		fmt.Println(number)
	}
}
