package watch

import (
	"fmt"
	"log"
	"time"

	"github.com/ondrejmalina/cli-watch/internal/cli"
	"github.com/ondrejmalina/cli-watch/internal/clock"
)

func Run(userInput cli.UserInput) {

	dur, err := time.ParseDuration(userInput.Argument)
	if err != nil {
		log.Fatalf("Cannot parse input time %v", userInput.Argument)
	}

	tick := time.Second
	ticker := time.NewTicker(tick)
	defer ticker.Stop()
	done := make(chan bool)
	go func() {
		time.Sleep(dur)
		done <- true
	}()
	for {
		select {
		case <-done:
			fmt.Println("Done!")
			return
		case <-ticker.C:
			cl := clock.CreateClock(dur)
			cl.Printf()
			dur -= tick
		}
	}
}
