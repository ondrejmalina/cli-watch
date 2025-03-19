package cli

import (
	"flag"
	"log"
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

	if len(flag.Args()) != 2 {
		log.Fatal("Wrong number of arguments, enter command and time")
	}

	command := flag.Arg(0)
	argument := flag.Arg(1)

	Inp := readInput(command, argument)
	return Inp
}
