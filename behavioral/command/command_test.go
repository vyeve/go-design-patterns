package command

func ExampleConsoleOutput() {
	queue := CommandQueue{}
	queue.AddCommand(CreateCommand("First message"))
	queue.AddCommand(CreateCommand("Second message"))
	queue.AddCommand(CreateCommand("Third message"))

	queue.AddCommand(CreateCommand("Fourth message"))
	queue.AddCommand(CreateCommand("Fifth message"))

	// Output:
	// Creating command
	// Creating command
	// Creating command
	// First message
	// Second message
	// Third message
	// Creating command
	// Creating command
}
