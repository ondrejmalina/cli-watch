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
	clock.StartTicking(tick)

	dur := time.Duration(0)
	cl := clock.Create(dur)
	fmt.Printf("\r%v", cl.Sprintf())

	for {
		select {
		case k := <-keyC:
			switch k {
			case 'p':
				clock.StopTicking()
			case 'r':
				clock.StartTicking(tick)
			case keyboard.KeyEsc, keyboard.KeyCtrlC:
				cl := clock.Create(dur)
				fmt.Printf("\rThe stopwatch was ticking for %v", cl.Sprintf())
				return
			}
		case <-clock.Ticker.C:
			dur += tick
			cl := clock.Create(dur)
			fmt.Printf("\r%v", cl.Sprintf())
		}
	}
}
