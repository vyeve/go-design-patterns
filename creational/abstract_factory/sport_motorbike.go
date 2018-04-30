package abstract_factory

type SportMotorbike struct{}

func (*SportMotorbike) NumWheels() int {
	return 2
}

func (*SportMotorbike) NumSeats() int {
	return 1
}

func (*SportMotorbike) GetMotorbikeType() int {
	return SportMotorbikeType
}
