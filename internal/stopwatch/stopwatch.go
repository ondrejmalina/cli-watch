package stopwatch

import (
	"fmt"
	"log"
	"time"

	"github.com/eiannone/keyboard"
	"github.com/ondrejmalina/cli-watch/internal/cli"
	"github.com/ondrejmalina/cli-watch/internal/clock"
)

func Run(userInput cli.UserInput) {

	dur, err := time.ParseDuration("99h")
	if err != nil {
		log.Fatalf("Cannot parse input time %v", userInput.Argument)
	}

	keyC := make(chan any, 1)
	go cli.KeyboardInput(keyC)

	tick := time.Second
	clock.StartTime(dur, tick)
	dur = 0

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
			case keyboard.KeyEsc, keyboard.KeyCtrlC:
				fmt.Printf("\rBye!")
				return
			}
		case <-clock.Ticker.C:
			dur += tick
			cl := clock.Create(dur)
			fmt.Printf("\r%v", cl.Sprintf())
		}
	}
}
