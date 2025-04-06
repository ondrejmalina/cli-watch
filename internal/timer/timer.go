package timer

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

	keyC := make(chan any, 1)
	go cli.KeyboardInput(keyC)

	tick := time.Second
	clk := clock.CreateClock()

	clk.StartTickerTimer(tick, dur)
	fmt.Printf("\r%v", clk.FmtDuration(dur))

	for {
		select {
		case k := <-keyC:
			switch k {
			case 'p':
				clk.StopTickerTimer()
			case 'r':
				clk.StartTickerTimer(tick, dur)
			case keyboard.KeyEsc, keyboard.KeyCtrlC:
				fmt.Printf("\r\033[KBye!\n")
				return
			}
		case <-clk.Ticker.C:
			dur -= tick
			fmt.Printf("\r%v", clk.FmtDuration(dur))
		case <-clk.Timer.C:
			dur -= tick
			fmt.Printf("\r%v", clk.FmtDuration(dur))
			fmt.Print("\r\033[KDone!\n")
			return

		}
	}
}
