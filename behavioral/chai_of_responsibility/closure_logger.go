package chor

type ClosureLogger struct {
	NextChain ChainLogger
	Closure   func(string)
}

func (c *ClosureLogger) Next(s string) {
	if c.Closure != nil {
		c.Closure(s)
	}

	if c.NextChain != nil {
		c.NextChain.Next(s)
	}
}
