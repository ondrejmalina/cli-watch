package clock

import (
	"fmt"
	"time"
)

type clock struct {
	hr, min, sec int64
}

func CreateClock(d time.Duration) clock {
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
