/*
Sometimes you may want to split values coming in from a channel so that you can send them off into two separate areas
of your codebase. Imagine a channel of user commands: you might want to take in a stream of user commands on a channel,
send them to something that executes them, and also send them to something that logs the commands for later auditing.

Taking its name from tee command in Unix-like systems.
*/

package tee_channel

func tee(done <-chan interface{}, in <-chan interface{}) (_, _ <-chan interface{}) {
	out1 := make(chan interface{})
	out2 := make(chan interface{})

	go func() {
		defer close(out1)
		defer close(out2)
		for val := range orDone(done, in) {
			// We will use local versions of out1 and out2, so we shadow these variables
			var out1, out2 = out1, out2
			// We are going to use one select statement so that writes to out1 and out2 don't block each other.
			// To ensure both are written to, we'll perform two iterations of the select statement: one for each
			// outbound channel.
			for i := 0; i < 2; i++ {
				select {
				case <-done:
				case out1 <- val:
					// Once we've written to a channel, we set its shadowed copy to nil so that further writes will
					// block and the other may continue.
					out1 = nil
				case out2 <- val:
					out2 = nil
				}
			}
		}
	}()
	return out1, out2
}

func orDone(done <-chan interface{}, c <-chan interface{}) <-chan interface{} {
	valStream := make(chan interface{})

	go func() {
		defer close(valStream)
		for {
			select {
			case <-done:
				return
			case v, ok := <-c:
				if !ok {
					return
				}
				select {
				case valStream <- v:
				case <-done:
					return
				}
			}
		}
	}()

	return valStream
}
