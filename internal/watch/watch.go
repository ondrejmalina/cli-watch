package watch

import (
	"fmt"
	"log"
	"time"

	"github.com/eiannone/keyboard"
	"github.com/ondrejmalina/cli-watch/internal/cli"
	"github.com/ondrejmalina/cli-watch/internal/clock"
)

func Run(userInput cli.UserInput) {

	dur, err := time.ParseDuration(userInput.Argument)
	if err != nil {
		log.Fatalf("Cannot parse input time %v", userInput.Argument)
	}

	keyC := make(chan rune, 1)

	// NOTE: Is this still running after pressing a key?
	// Why is the for loop necessary?
	go func() {
		for {
			r, _, err := keyboard.GetSingleKey()
			if err != nil {
				log.Fatal("Invalid keyboard input")
			}
			keyC <- r
		}
	}()

	tick := time.Second
	clock.StartTime(dur, tick)

	cl := clock.Create(dur)
	fmt.Printf("\r%v", cl.Sprintf())

	for {
		select {
		case k := <-keyC:
			switch k {
			case 'p':
				clock.StopTime()
			case 'r':
				clock.StartTime(dur, tick)
			default:
				fmt.Print(string(k))
			}
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
