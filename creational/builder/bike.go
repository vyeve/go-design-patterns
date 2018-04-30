package builder

type BikeBuilder struct {
	v VehicleProduct
}

func (bb *BikeBuilder) SetWheels() BuildProcess {
	bb.v.Wheels = 2
	return bb
}

func (bb *BikeBuilder) SetSeats() BuildProcess {
	bb.v.Seats = 2
	return bb
}

func (bb *BikeBuilder) SetStructure() BuildProcess {
	bb.v.Structure = "Motorbike"
	return bb
}

func (bb *BikeBuilder) GetVehicle() VehicleProduct {
	return bb.v
}
