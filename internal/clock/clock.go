package clock

import (
	"fmt"
	"time"
)

var (
	Timer  *time.Timer
	Ticker *time.Ticker
)

func StartTime(dur time.Duration, tick time.Duration) {
	Timer = time.NewTimer(dur)
	Ticker = time.NewTicker(tick)
}

func StopTime() {
	Timer.Stop()
	Ticker.Stop()
}

func StartTicking(tick time.Duration) {
	Ticker = time.NewTicker(tick)
}

func StopTicking() {
	Ticker.Stop()
}

type clock struct {
	hr, min, sec int64
}

// TODO: Rename
func Create(d time.Duration) clock {
	d = d.Round(time.Second)
	hr := d / time.Hour
	d -= hr * time.Hour
	min := d / time.Minute
	d -= min * time.Minute
	sec := d / time.Second
	d -= min * time.Second
	return clock{int64(hr), int64(min), int64(sec)}
}

func (t *clock) Sprintf() string {
	return fmt.Sprintf("%02d:%02d:%02d", t.hr, t.min, t.sec)
}
