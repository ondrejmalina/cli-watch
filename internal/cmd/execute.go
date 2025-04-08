package cmd

import (
	"fmt"

	"github.com/ondrejmalina/cli-watch/internal/cli"
	"github.com/ondrejmalina/cli-watch/internal/stopwatch"
	"github.com/ondrejmalina/cli-watch/internal/timer"
)

func Execute() error {

	userInput := cli.ParseInput()

	switch userInput.Command {
	case "timer":
		timer.Run(userInput)
	case "stopwatch":
		stopwatch.Run(userInput)
	default:
		return fmt.Errorf("Unknown command %v", userInput.Command)
	}
	return nil
}
