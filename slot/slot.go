package slot

import (
	"errors"
	"time"
)

// Slot structure
type Slot struct {
	slotID        string
	VehicleNumber string
	entryTime     time.Time
	exitTime      time.Time
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
	entryTime := time.Now()
	s.VehicleNumber = vNumber
	s.entryTime = entryTime
}

// CarExit method will set vehicle detailsS
func (s *Slot) CarExit() {
	exitTime := time.Now()
	s.exitTime = exitTime
	diff := s.exitTime.Sub(s.entryTime)
	s.pay = s.slotPrice * float32(diff)
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
	entryTime := time.Now()
	s.VehicleNumber = vNumber
	s.entryTime = entryTime
}

// VanExit method will set vehicle detailsS
func (s *Slot) VanExit() {
	exitTime := time.Now()
	s.exitTime = exitTime
	diff := s.exitTime.Sub(s.entryTime)
	s.pay = s.slotPrice * float32(diff)
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
	entryTime := time.Now()
	s.VehicleNumber = vNumber
	s.entryTime = entryTime
}

// BikeExit method will set vehicle detailsS
func (s *Slot) BikeExit() {
	exitTime := time.Now()
	s.exitTime = exitTime
	diff := s.exitTime.Sub(s.entryTime)
	s.pay = s.slotPrice * float32(diff)
}
