package composite

func ExampleAthleteA() {
	swimmer := CompositeSwimmerA{MySwim: Swim}

	swimmer.MyAthlete.Train()
	swimmer.MySwim()
	// Output: Training
	// Swimming!
}

func ExampleAthleteB() {
	swimmer := CompositeSwimmerB{
		Trainer: new(Athlete),
		Swimmer: new(SwimmerImpl),
	}
	swimmer.Train()
	swimmer.Swim()
	// Output: Training
	// Swimming!
}
