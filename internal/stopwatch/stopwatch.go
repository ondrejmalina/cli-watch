package stopwatch

import (
	"fmt"
	"time"

	"github.com/eiannone/keyboard"
	"github.com/ondrejmalina/cli-watch/internal/cli"
	"github.com/ondrejmalina/cli-watch/internal/clock"
)

func Run(userInput cli.UserInput) {

	keyC := make(chan any, 1)
	go cli.KeyboardInput(keyC)

	tick := time.Second
	clk := clock.CreateClock()

	dur := time.Duration(0)
	clk.StartTicker(tick)
	fmt.Printf("\r%v", clk.FmtDuration(dur))

	for {
		select {
		case k := <-keyC:
			switch k {
			case 'p':
				clk.StopTicker()
			case 'r':
				clk.StartTicker(tick)
			case keyboard.KeyEsc, keyboard.KeyCtrlC:
				fmt.Printf("\rThe stopwatch was ticking for %v",
					clk.FmtDuration(dur))
				return
			}
		case <-clk.Ticker.C:
			dur += tick
			fmt.Printf("\r%v", clk.FmtDuration(dur))
		}
	}
}
