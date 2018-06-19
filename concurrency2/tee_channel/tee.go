/*
Sometimes you may want to split values coming in from a channel so that you can send them off into two separate areas
of your codebase. Imagine a channel of user commands: you might want to take in a stream of user commands on a channel,
send them to something that executes them, and also send them to something that logs the commands for later auditing.

Taking its name from tee command in Unix-like systems.
*/

package tee_channel

import ""

func tee(done <-chan interface{}, in <-chan interface{}) (_, _ <-chan interface{}) {
	out1 := make(chan interface{})
	out2 := make(chan interface{})

	go func() {
		defer close(out1)
		defer close(out2)
		for val := range or_done_channel.orDone(done, in) {
			var out1, out2 = out1, out2
			for i := 0; i < 2; i++ {
				select {
				case <-done:
				case out1 <- val:
					out1 = nil
				case out2 <- val:
					out2 = nil
				}
			}
		}
	}()
	return out1, out2
}
