/*
The command design pattern focus on the invocation of something or on the abstraction of some type.
A Command pattern is commonly seen as a container. The Command pattern will be used heavily when
dealing with channels. With channels you can send any message through it but, if we need a response
from the receiving side of the channel, a common approach is to create a command that has a second,
response channel attached where we are listening.
*/

package command

type Command interface {
	Execute()
}

type CommandQueue struct {
	queue []Command
}

func (q *CommandQueue) AddCommand(c Command) {
	q.queue = append(q.queue, c)
	if len(q.queue) == 3 {
		for _, command := range q.queue {
			command.Execute()
		}
		q.queue = make([]Command, 0, 3)
	}
}
