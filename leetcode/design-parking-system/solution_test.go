package design_parking_system

import "testing"

func TestParkingSystem_AddCar(t *testing.T) {
	obj := Constructor(1, 1, 0)
	if !obj.AddCar(1) {
		t.Fatal("failed")
	}
	if !obj.AddCar(2) {
		t.Fatal("failed")
	}
	if obj.AddCar(3) {
		t.Fatal("failed")
	}
	if obj.AddCar(1) {
		t.Fatal("failed")
	}
}
