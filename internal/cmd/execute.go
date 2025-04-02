package cmd

import (
	"log"

	"github.com/ondrejmalina/cli-watch/internal/cli"
	"github.com/ondrejmalina/cli-watch/internal/stopwatch"
	"github.com/ondrejmalina/cli-watch/internal/watch"
)

// Main driving function of the module.
func Execute() {

	userInput := cli.ParseInput()

	switch userInput.Command {
	case "watch":
		watch.Run(userInput)
	case "stopwatch":
		stopwatch.Run(userInput)
	default:
		log.Fatalf("Unknown command %v", userInput.Command)
	}
}
