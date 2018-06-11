The Barrier pattern puts up a barrier so that nobody passes until we have all the results we need.

Imagine the situation where we have a microservices application where one service needs to compose its response by 
merging the responses of another three microservices. 

Our Barrier pattern could be a service that will block its response until it has been composed with the results
returned by one or more different Goroutines (or services).

The Barrier pattern is not only useful to make network requests; it could be also used to split some task into
multiple Goroutines. For example, an expensive operation could be split into a few smaller operations distributed in 
different Goroutines to maximize parallelism and achieve better performance.