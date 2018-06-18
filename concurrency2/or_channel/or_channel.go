/*
At times you may find yourself wanting to combine one or more done channels into a single done channel that closes if
any of its component channels close. It is perfectly acceptable, albeit verbose, to write a select statement that
performs this coupling; however, sometimes you can't know the number of done channels you're working with at runtime.
In this case, or if you just prefer a one-liner, you can combine these channels together using the or-channel pattern.
*/

package or_channel

func OrChannel(channels ...<-chan interface{}) <-chan interface{} {
	switch len(channels) {
	case 0:
		return nil
	case 1:
		return channels[0]
	}

	orDone := make(chan interface{})

	go func() {
		defer close(orDone)

		switch len(channels) {
		case 2:
			select {
			case <-channels[0]:
			case <-channels[1]:
			}
		default:
			select {
			case <-channels[0]:
			case <-channels[1]:
			case <-channels[2]:
			case <-OrChannel(append(channels[3:], orDone)...):
			}
		}
	}()
	return orDone
}
