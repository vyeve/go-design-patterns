In some circumstances, you may find yourself wanting to consume values from a sequence of channels:
```go
<-chan <-chan interface{}
```

A sequence of channels suggests an ordered write, albeit from different sources.

As a consumer, the code may not care about the fact that its values come from a sequence of channels. In that case, 
dealing with a channel of channels can be cumbersome. If we instead define a function that can destructure the channel 
of channels into a simple channel — a technique called *bridging* the channels — this will make it much easier for the 
consumer to focus on the problem at hand.