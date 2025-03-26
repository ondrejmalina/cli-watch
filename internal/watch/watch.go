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
	clock.StartTime(dur, tick)

	cl := clock.Create(dur)
	fmt.Printf("\r%v", cl.Sprintf())

	for {
		select {
		case <-clock.Ticker.C:
			dur -= tick
			// NOTE: Put into goroutine?
			cl := clock.Create(dur)
			fmt.Printf("\r%v", cl.Sprintf())
		case <-clock.Timer.C:
			dur -= tick
			cl := clock.Create(dur)
			fmt.Printf("\r%v", cl.Sprintf())
			return

		}
	}
}
