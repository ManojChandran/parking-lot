package parking

import (
	"errors"
	"fmt"

	"github.com/ManojChandran/parking-lot/slot"
	"github.com/google/uuid"
)

// Ticket delivers the parking ticket for vehicle
type Ticket struct {
	ID string
}

type pEntrance struct {
	gateLocation string
	gateNumber   string
}

type parking struct {
	parkingCompany string
	parkingAddress string
	//	parkingEntrance [3]pEntrance
	carParking  map[Ticket]*slot.Slot
	vanParking  map[Ticket]*slot.Slot
	bikeParking map[Ticket]*slot.Slot
}

// Parking interface provide the list of avialable methods
type Parking interface {
	SetParkingCompany(string, string)
	GetParkingTicket(string, string) (Ticket, error)
	GetVehicle(string) string
	RemoveVehicle(string)
	DisplayParking()
}

func (p *parking) SetParkingCompany(parkingCompany, parkingAddress string) {
	p.parkingCompany = parkingCompany
	p.parkingAddress = parkingAddress
}

// BuildParkingCompany method will create a parking company instance
func BuildParkingCompany() Parking {
	CarParking := make(map[Ticket]*slot.Slot)
	VanParking := make(map[Ticket]*slot.Slot)
	BikeParking := make(map[Ticket]*slot.Slot)
	return &parking{
		carParking:  CarParking,
		vanParking:  VanParking,
		bikeParking: BikeParking,
	}
}

func createTicketID() Ticket {
	ID := uuid.New()
	return Ticket{ID: ID.String()}
}

// GetParkingTicket builds parking ticket
func (p *parking) GetParkingTicket(vType, vNumber string) (Ticket, error) {
	var t Ticket
	if vType == "car" {
		c, err := slot.NewCarSlot()
		if err != nil {
			return t, errors.New("no slot available")
		}
		c.SetCarSlot(vNumber)
		t = createTicketID()
		p.carParking[t] = c
	}

	if vType == "van" {
		v, err := slot.NewVanSlot()
		if err != nil {
			return t, errors.New("no slot available")
		}
		v.SetVanSlot(vNumber)
		t = createTicketID()
		p.vanParking[t] = v
	}

	if vType == "bike" {
		b, err := slot.NewBikeSlot()
		if err != nil {
			return t, errors.New("no slot available")
		}
		b.SetBikeSlot(vNumber)
		t = createTicketID()
		p.bikeParking[t] = b
	}

	return t, nil
}

func (p *parking) GetVehicle(id string) string {
	t := Ticket{ID: id}
	pop, ok := p.carParking[t]
	if ok {
		return pop.VehicleNumber
	}
	return "no vehicle"
}

func (p *parking) RemoveVehicle(id string) {
	t := Ticket{ID: id}
	delete(p.carParking, t)
	delete(p.vanParking, t)
	delete(p.bikeParking, t)
}

func (p *parking) DisplayParking() {

	fmt.Println(p.parkingCompany)
	fmt.Println(p.parkingAddress)
	fmt.Println("*****Car *****")
	fmt.Println(slot.CarCount())
	for _, c := range p.carParking {
		fmt.Println(c)
	}
	fmt.Println("*****Van *****")
	fmt.Println(slot.VanCount())
	for _, v := range p.vanParking {
		fmt.Println(v)
	}
	fmt.Println("*****Bike ****")
	fmt.Println(slot.BikeCount())
	for _, b := range p.bikeParking {
		fmt.Println(b)
	}
	//	fmt.Println(p.parkingEntrance)
	// fmt.Println(p.parkingFloor)

}
