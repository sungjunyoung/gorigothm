package design_parking_system

type ParkingSystem struct {
	spaces  [4]int
	current [4]int
}

func Constructor(big int, medium int, small int) ParkingSystem {
	return ParkingSystem{
		spaces:  [4]int{-1, big, medium, small},
		current: [4]int{-1, 0, 0, 0},
	}
}

func (this *ParkingSystem) AddCar(carType int) bool {
	if this.current[carType] == this.spaces[carType] {
		return false
	}

	this.current[carType] += 1
	return true
}

/**
 * Your ParkingSystem object will be instantiated and called as such:
 * obj := Constructor(big, medium, small);
 * param_1 := obj.AddCar(carType);
 */
