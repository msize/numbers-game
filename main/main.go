package main

import (
	"os"

	"github.com/numbers-game/game/human_vs_human"
	"github.com/numbers-game/game/human_vs_machine"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		return
	}
	switch args[0] {
	case "hh":
		human_vs_human.Run()
	case "hm":
		human_vs_machine.Run()
	}
}
