package pipeline

func LaunchPipeline(amount int) int {
	return <-sum(power(generator(amount)))
}

func generator(max int) <-chan int {
	out := make(chan int, max)

	go func() {
		for i := 1; i <= max; i++ {
			out <- i
		}
		close(out)
	}()

	return out
}

func power(in <-chan int) <-chan int {
	out := make(chan int, cap(in))

	go func() {
		for v := range in {
			out <- v * v
		}
		close(out)
	}()
	return out
}

func sum(in <-chan int) <-chan int {
	out := make(chan int)

	go func() {
		sum := 0
		for v := range in {
			sum += v
		}
		out <- sum
		close(out)
	}()
	return out
}
