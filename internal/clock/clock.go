package clock

import (
	"fmt"
	"time"
)

type Clock struct {
	Timer  *time.Timer
	Ticker *time.Ticker
}

func CreateClock() *Clock {
	return &Clock{}
}

func (c *Clock) StartTicker(tickDuration time.Duration) {
	c.Ticker = time.NewTicker(tickDuration)
}

func (c *Clock) StopTicker() {
	c.Ticker.Stop()
}

func (c *Clock) StartTickerTimer(tickDuration time.Duration, timeDuration time.Duration) {
	c.Timer = time.NewTimer(timeDuration)
	c.Ticker = time.NewTicker(tickDuration)
}

func (c *Clock) StopTickerTimer() {
	if c.Ticker != nil {
		c.Ticker.Stop()
	}
	if c.Timer != nil {
		if !c.Timer.Stop() {
			select {
			case <-c.Timer.C: // Drain the channel if timer already fired
			default:
			}
		}
	}
}

func (c *Clock) FmtDuration(d time.Duration) string {
	d = d.Round(time.Second)
	hr := d / time.Hour
	d -= hr * time.Hour
	min := d / time.Minute
	d -= min * time.Minute
	sec := d / time.Second
	d -= min * time.Second
	return fmt.Sprintf("%02d:%02d:%02d", hr, min, sec)
}
