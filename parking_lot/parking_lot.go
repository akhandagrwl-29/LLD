package parkinglot

import (
	"errors"
	"fmt"
)

var instance *ParkingLot

type ParkingLot struct {
	levels []*Level
}

func NewParkingLot() *ParkingLot {
	if instance != nil {
		return instance
	}
	return &ParkingLot{
		levels: []*Level{},
	}
}

func (p *ParkingLot) AddLevel(level *Level) error {
	for _, currLevel := range p.levels {
		if currLevel.floor == level.floor {
			return errors.New("level already exists")
		}
	}
	p.levels = append(p.levels, level)
	return nil
}

func (p *ParkingLot) ShowAvailability() {
	fmt.Println("Showing Availability")
	for idx, level := range p.levels {
		fmt.Printf("--- Level: %d Floor: %d------\n", idx+1, level.floor)
		for _, spot := range level.parkingSpots {
			if spot.IsOccupied() {
				fmt.Printf("spot: %d is occupied with: %+v\n", spot.GetParkingSpot().id, VehicleType(spot.GetParkingSpot().vehileType))
			} else {
				fmt.Printf("Spot: %d is empty of vehicle type : %+v\n", spot.GetParkingSpot().id, VehicleType(spot.GetParkingSpot().vehileType))
			}
		}
	}
}

func (p *ParkingLot) ParkVehicle(vehicle Vehicle) bool {
	for _, level := range p.levels {
		if level.ParkVehicle(vehicle) {
			return true
		}
	}
	return false
}

func (p *ParkingLot) UnparkVehicle(vehicle Vehicle) bool {
	for _, level := range p.levels {
		if level.UnParkVehicle(vehicle) {
			return true
		}
	}
	return false
}
