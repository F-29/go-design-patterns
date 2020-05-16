package builder

import (
	"reflect"
	"testing"
)

var ManufacturingComplex = ManufacturingDirector{}

/***/

func carTestSetupHelper() VehicleProduct {
	carBuilder := &CarBuilder{}
	ManufacturingComplex.SetBuilder(carBuilder)
	ManufacturingComplex.Construct()

	return carBuilder.GetVehicle()
}

func bikeTestSetupHelper() VehicleProduct {
	bikeBuilder := &BikeBuilder{}
	ManufacturingComplex.SetBuilder(bikeBuilder)
	ManufacturingComplex.Construct()

	return bikeBuilder.GetVehicle()
}

func busTestSetupHelper() VehicleProduct {
	busBuilder := &BusBuilder{}
	ManufacturingComplex.SetBuilder(busBuilder)
	ManufacturingComplex.Construct()

	return busBuilder.GetVehicle()
}

/***/

func TestCarBuilder_SetSeats(t *testing.T) {
	var car = carTestSetupHelper()
	if car.Seats != 5 {
		t.Errorf("Expected Seats in car to be: 5 got %+v\n", car.Seats)
	}
}

func TestCarBuilder_SetWheels(t *testing.T) {
	var car = carTestSetupHelper()
	if car.Wheels != 4 {
		t.Errorf("Expected Wheels in car to be: 4 got %+v\n", car.Wheels)
	}
}

func TestCarBuilder_SetStructure(t *testing.T) {
	var car = carTestSetupHelper()

	if car.Structure != "Car" {
		t.Errorf("Expected Structure of car to be: 'Car' got %+v\n", car.Structure)
	}
}

func TestCarBuilder_GetVehicle(t *testing.T) {
	var car = carTestSetupHelper()
	if reflect.TypeOf(car) != reflect.TypeOf(VehicleProduct{}) {
		t.Errorf("Expected type of car to be: 'VehicleProduct' got %+v\n", reflect.TypeOf(car))
	}
}

/***/

func TestBikeBuilder_SetWheels(t *testing.T) {
	bike := bikeTestSetupHelper()
	if bike.Wheels != 2 {
		t.Errorf("Expected Wheels in bike to be: 2 got %+v\n", bike.Wheels)
	}
}

func TestBikeBuilder_SetSeats(t *testing.T) {
	bike := bikeTestSetupHelper()
	if bike.Seats != 2 {
		t.Errorf("Expected Seats in bike to be: 2 got %+v\n", bike.Seats)
	}
}

func TestBikeBuilder_SetStructure(t *testing.T) {
	bike := bikeTestSetupHelper()
	if bike.Structure != "Motorbike" {
		t.Errorf("Expected Structure of bike to be: 'Motorbike' got %+v\n", bike.Structure)
	}
}

func TestBikeBuilder_GetVehicle(t *testing.T) {
	bike := bikeTestSetupHelper()
	if reflect.TypeOf(bike) != reflect.TypeOf(VehicleProduct{}) {
		t.Errorf("Expected type of bike to be: 'builder.VehicleProduct' got %+v\n", reflect.TypeOf(bike))
	}

}

/***/

func TestBusBuilder_SetSeats(t *testing.T) {
	var bus = busTestSetupHelper()
	if bus.Seats != 30 {
		t.Errorf("Expected Seats in bus to be: 30 got %+v\n", bus.Seats)
	}
}

func TestBusBuilder_SetWheels(t *testing.T) {
	var bus = busTestSetupHelper()
	if bus.Wheels != 8 {
		t.Errorf("Expected Wheels in bus to be: 8 got %+v\n", bus.Wheels)
	}
}

func TestBusBuilder_SetStructure(t *testing.T) {
	var bus = busTestSetupHelper()
	if bus.Structure != "Bus" {
		t.Errorf("Expected Structure of bus to be: 'Bus' got %+v\n", bus.Structure)
	}
}

func TestBusBuilder_GetVehicle(t *testing.T) {
	bus := busTestSetupHelper()
	if reflect.TypeOf(bus) != reflect.TypeOf(VehicleProduct{}) {
		t.Errorf("Expected type of bus to be: 'builder.VehicleProduct' got %+v\n", reflect.TypeOf(bus))
	}

}
