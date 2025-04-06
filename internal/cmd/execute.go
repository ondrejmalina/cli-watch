package cmd

import (
	"fmt"

	"github.com/ondrejmalina/cli-watch/internal/cli"
	"github.com/ondrejmalina/cli-watch/internal/stopwatch"
	"github.com/ondrejmalina/cli-watch/internal/watch"
)

// Main driving function of the module.
func Execute() error {

	userInput := cli.ParseInput()

	switch userInput.Command {
	case "watch":
		watch.Run(userInput)
	case "stopwatch":
		stopwatch.Run(userInput)
	default:
		return fmt.Errorf("Unknown command %v", userInput.Command)
	}
	return nil
}
