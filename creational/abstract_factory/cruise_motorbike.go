package abstract_factory

type CruiseMotorbike struct{}

func (*CruiseMotorbike) NumWheels() int {
	return 2
}

func (*CruiseMotorbike) NumSeats() int {
	return 2
}

func (*CruiseMotorbike) GetMotorbikeType() int {
	return CruiseMotorbikeType
}
