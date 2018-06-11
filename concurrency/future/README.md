The Future design pattern (also called **Promise**) is a quick and easy way to achieve concurrent structures for 
asynchronous programming.

In short, we will define each possible behavior of an action before executing them in different Goroutines. The idea
here is to achieve a *fire-and-forget* that handles all possible results in an action.