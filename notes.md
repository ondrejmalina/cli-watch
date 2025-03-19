Check https://github.com/antonmedv/countdown/blob/master/main.go
- `timeLeft` is Duration from cli input that enter `countdown` function.
- `countdown` function triggers `start` function which start ticker(`tick` which defaults to second) and timer(`timeLeft`).
- the `countdown` fn subtracts ticks from the `timeLeft` duration. On each tick, it runs a `draw` fn, which executes `format` fn, which handles the transformation from `Duration` format into a nice string.
