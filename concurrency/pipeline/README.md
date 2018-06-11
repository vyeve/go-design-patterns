Every time that we write any function that performs some logic, we are writing a pipeline. If *this* then *that*, or
else *something else*. Pipelines pattern can be made more complex by using a few functions that call to each other.
They can even get looped in their out execution.

The Pipeline pattern in Go work in a similar fashion, but each step in the Pipeline will be in a different Goroutine and
communication, and synchronizing will be done using channels.

```go
package pipeline

func functionName (in <-chan int) <-chan int {
	out := make(chan int, 100)
	
	go func() {
		for v := range in {
			// Do something with v and send it to channel out
		}
		close(out)
	}()
	
	return out
}

```