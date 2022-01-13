package slot

import (
	"errors"
	"time"
)

// Slot structure
type Slot struct {
	slotID        string
	VehicleNumber string
	entryTime     string
	exitTime      string
	slotPrice     float32
	pay           float32
}

var carInstanceCnt int
var vanInstanceCnt int
var bikeInstanceCnt int

// NewCarSlot will construct slot
func NewCarSlot() (*Slot, error) {
	if carInstanceCnt > 5 {
		return nil, errors.New("no slot available")
	}
	carInstanceCnt++
	return &Slot{slotPrice: 10.00}, nil
}

// CarCount returns no of car
func CarCount() int {
	return carInstanceCnt
}

// SetCarSlot method will set vehicle detailsS
func (s *Slot) SetCarSlot(vNumber string) {
	entryTime := time.Now().Format(time.RFC850)
	s.VehicleNumber = vNumber
	s.entryTime = entryTime
}

// CarExitTime method will set vehicle detailsS
func (s *Slot) CarExitTime() {
	exitTime := time.Now().Format(time.RFC850)
	s.exitTime = exitTime
}

// NewVanSlot will construct slot
func NewVanSlot() (*Slot, error) {
	if vanInstanceCnt > 3 {
		return nil, errors.New("no slot available")
	}
	vanInstanceCnt++
	return &Slot{slotPrice: 15.00}, nil
}

// VanCount returns no of van
func VanCount() int {
	return vanInstanceCnt
}

// SetVanSlot method will set vehicle detailsS
func (s *Slot) SetVanSlot(vNumber string) {
	entryTime := time.Now().Format(time.RFC850)
	s.VehicleNumber = vNumber
	s.entryTime = entryTime
}

// VanExitTime method will set vehicle detailsS
func (s *Slot) VanExitTime() {
	exitTime := time.Now().Format(time.RFC850)
	s.exitTime = exitTime
}

// NewBikeSlot will construct slot
func NewBikeSlot() (*Slot, error) {
	if bikeInstanceCnt > 2 {
		return nil, errors.New("no slot available")
	}
	bikeInstanceCnt++
	return &Slot{slotPrice: 5.00}, nil
}

// BikeCount returns no of bike
func BikeCount() int {
	return bikeInstanceCnt
}

// SetBikeSlot method will set vehicle detailsS
func (s *Slot) SetBikeSlot(vNumber string) {
	entryTime := time.Now().Format(time.RFC850)
	s.VehicleNumber = vNumber
	s.entryTime = entryTime
}

// BikeExitTime method will set vehicle detailsS
func (s *Slot) BikeExitTime() {
	exitTime := time.Now().Format(time.RFC850)
	s.exitTime = exitTime
}
