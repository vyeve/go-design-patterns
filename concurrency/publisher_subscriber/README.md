The *Observer* pattern maintains a list of observers or subscribers that want to be notified of a particular event. In 
this case, each subscriber is going to run in a different Goroutine as well as the publisher. We will have new problems
with building this structure:
* Now, the access to the list of subscriber must be serialized. If we are reading the list with one Goroutine, we cannot
be removing a subscriber from it or we will have a race.
* When a subscriber is removed, the subscriber's Goroutine must be closed too, or it will keep iterating forever and we
will run into Goroutine leaks.
* When stopping the publisher, all subscribers must stop their Goroutines, too.