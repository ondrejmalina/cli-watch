package main

import (
	"fmt"
	"os"

	cmd "github.com/ondrejmalina/cli-watch/internal/cmd"
)

func main() {
	err := cmd.Execute()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
