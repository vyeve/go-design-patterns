package chor

import (
	"fmt"
	"strings"
)

type SecondLogger struct {
	NextChain ChainLogger
}

func (l *SecondLogger) Next(s string) {
	if strings.Contains(strings.ToLower(s), "hello") {
		fmt.Printf("Second logger: %s\n", s)
		if l.NextChain != nil {
			l.NextChain.Next(s)
		}
		return
	}
	fmt.Printf("Finishing in second logging\n\n")
}
