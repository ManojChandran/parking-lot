package main

import (
	"fmt"

	"github.com/ManojChandran/parking-lot/parking"
)

func main() {
	p := parking.BuildParkingCompany()
	p.SetParkingCompany("Zam Parking", "Connecticut")
	t1, err := p.GetParkingTicket("car", "KL BT 6986")
	fmt.Println(t1.ID, err)
	t2, err := p.GetParkingTicket("car", "KL BT 6888")
	fmt.Println(t2.ID, err)
	t3, err := p.GetParkingTicket("car", "KL BT 6899")
	fmt.Println(t3.ID, err)
	t4, err := p.GetParkingTicket("car", "KL BT 6899")
	fmt.Println(t4.ID, err)
	t5, err := p.GetParkingTicket("car", "KL BT 6899")
	fmt.Println(t5.ID, err)
	t6, err := p.GetParkingTicket("car", "KL BT 6899")
	fmt.Println(t6.ID, err)
	t7, err := p.GetParkingTicket("car", "KL BT 6899")
	fmt.Println(t7.ID, err)
	v1, err := p.GetParkingTicket("van", "KL BT 6986")
	fmt.Println(v1.ID, err)
	v2, err := p.GetParkingTicket("van", "KL BT 6986")
	fmt.Println(v2.ID, err)
	b1, err := p.GetParkingTicket("bike", "KL BT 6986")
	fmt.Println(b1.ID, err)
	b2, err := p.GetParkingTicket("bike", "KL BT 6986")
	fmt.Println(b2.ID, err)
	v := p.GetVehicle(t1.ID)
	u := p.GetVehicle(t2.ID)
	fmt.Println(v, u)
	p.DisplayParking()
	p.VehicleExit(t2.ID)
	p.VehicleExit(v1.ID)
	p.VehicleExit(b2.ID)
	p.DisplayParking()

}
