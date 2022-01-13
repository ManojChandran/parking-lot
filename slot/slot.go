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
	return &Slot{}, nil
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

// NewVanSlot will construct slot
func NewVanSlot() (*Slot, error) {
	if vanInstanceCnt < 3 {
		return nil, errors.New("no slot available")
	}
	vanInstanceCnt++
	return &Slot{}, nil
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

// NewBikeSlot will construct slot
func NewBikeSlot() (*Slot, error) {
	if bikeInstanceCnt < 2 {
		return nil, errors.New("no slot available")
	}
	bikeInstanceCnt++
	return &Slot{}, nil
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
