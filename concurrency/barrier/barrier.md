The Barrier pattern puts up a barrier so that nobody passes until we have all the results we need.

Imagine the situation where we have a microservices application where one service needs to compose its response by 
merging the responses of another three microservices. 

Our Barrier pattern could be a service that will block its response until it has been composed with the results
returned by one or more different Goroutines (or services).