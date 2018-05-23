package chor

import "fmt"

type FirstLogger struct {
	NextChain ChainLogger
}

func (l *FirstLogger) Next(s string) {
	fmt.Printf("First logger: %s\n", s)
	if l.NextChain != nil {
		l.NextChain.Next(s)
	}
}
