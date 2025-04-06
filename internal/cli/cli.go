package cli

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/eiannone/keyboard"
)

func printHelp() {
	help := `
Usage: cli-watch [command] [golang time fmt]

Commands:
  - timer [golang time fmt]: start a timer
  - stopwatch: start a stopwatch

Controls:
  - p: pause running watch
  - r: resume paused watch
  - esc: stop watch

Examples:
  - cli-watch timer 2m30s
  - cli-watch stopwatch
`
	fmt.Print(help)
}

type UserInput struct {
	Command, Argument string
}

func ParseInput() UserInput {
	flag.Usage = printHelp
	flag.Parse()

	cmd := flag.Arg(0)
	if cmd == "" {
		printHelp()
		os.Exit(0)
	}

	return UserInput{
		Command:  flag.Arg(0),
		Argument: flag.Arg(1),
	}
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
