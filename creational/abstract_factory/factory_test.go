package abstract_factory

import "testing"

func TestMotorbikeFactory(t *testing.T) {
	motorbikeF, err := BuildFactory(MotorbikeFactoryType)
	if err != nil {
		t.Fatal(err)
	}

	motorbikeVehicle, err := motorbikeF.NewVehicle(SportMotorbikeType)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("motorbike vehicle has %d wheels", motorbikeVehicle.NumWheels())

	sportbike, ok := motorbikeVehicle.(Motorbike)
	if !ok {
		t.Fatal("struct assertion has failed")
	}
	t.Logf("Sport motorbike has type %d\n", sportbike.GetMotorbikeType())
}

func TestCarFactory(t *testing.T) {
	carF, err := BuildFactory(CarFactoryType)
	if err != nil {
		t.Fatal(err)
	}
	carVehicle, err := carF.NewVehicle(LuxuryCarType)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Car vehicle has %d seats\n", carVehicle.NumSeats())
	luxuryCar, ok := carVehicle.(Car)
	if !ok {
		t.Fatal("struct assertion has failed")
	}
	t.Logf("Luxury car has %d doors.\n", luxuryCar.NumDoors())
}
