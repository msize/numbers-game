package main

import (
	"fmt"

	"github.com/numbers-game/permutations"
)

func main() {
	numbers := permutations.New()
	for i := 0; i < numbers.Len(); i++ {
		fmt.Println(numbers.Get(i))
	}
}
