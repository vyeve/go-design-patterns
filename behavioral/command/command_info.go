package command

import "time"

type Commander interface {
	Info() string
}

type TimePassed struct {
	Start time.Time
}

func (t *TimePassed) Info() string {
	return time.Since(t.Start).String()
}

type HelloMessage struct{}

func (*HelloMessage) Info() string {
	return "Hello world!"
}
