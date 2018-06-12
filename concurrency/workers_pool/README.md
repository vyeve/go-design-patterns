With a pool of workers, we want to bound the amount of Goroutines available so that we have a deeper control of the pool
of resources. This is easy to achieve by creating a channel for each worker and having workers with either an idle or 
busy status.

The workers pool design pattern helps us to do the following:
* Control access to shared resources using quotas
* Create a limited amount of Goroutines per app
* Provide more parallelism capabilities to other concurrent structures