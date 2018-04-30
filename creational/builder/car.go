package builder

type CarBuilder struct {
	v VehicleProduct
}

func (cb *CarBuilder) SetWheels() BuildProcess {
	cb.v.Wheels = 4
	return cb
}

func (cb *CarBuilder) SetSeats() BuildProcess {
	cb.v.Seats = 5
	return cb
}

func (cb *CarBuilder) SetStructure() BuildProcess {
	cb.v.Structure = "Car"
	return cb
}

func (cb *CarBuilder) GetVehicle() VehicleProduct {
	return cb.v
}
