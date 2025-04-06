package cli

import (
	"flag"
	"log"

	"github.com/eiannone/keyboard"
)

type UserInput struct {
	Command, Argument string
}

func readInput(cmd string, arg string) UserInput {
	inp := UserInput{cmd, arg}
	return inp
}

func ParseInput() UserInput {

	flag.Parse()

	command := flag.Arg(0)
	argument := flag.Arg(1)

	input := readInput(command, argument)
	return input
}

func KeyboardInput(keyC chan any) {

	for {
		r, k, err := keyboard.GetSingleKey()
		if err != nil {
			log.Fatalf("Error reading keyboard input: %v", err)
		}
		switch k {
		case 0:
			keyC <- r
		default:
			keyC <- k
		}
	}
}
