package or_done_channel

func orDone(done <-chan interface{}, c <-chan struct{}) <-chan struct{} {
	valStream := make(chan struct{})

	go func() {
		defer close(valStream)
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
	}()

	return valStream
}

/*
Example:

for val := range orDone(done, myChan) {
	// Do something with val
}
*/
