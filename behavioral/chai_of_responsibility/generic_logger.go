package chor

import (
	"fmt"
	"io"
)

type WriterLogger struct {
	NextChain ChainLogger
	Writer    io.Writer
}

func (l *WriterLogger) Next(s string) {
	if l.Writer != nil {
		l.Writer.Write([]byte(fmt.Sprintf("WriterLogger: %s\n", s)))
	}
	if l.NextChain != nil {
		l.NextChain.Next(s)
	}
}
