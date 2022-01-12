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
	v := p.GetVehicle(t1.ID)
	u := p.GetVehicle(t2.ID)
	/*	p.AddParkingEntrance("North Gate", "no.1")
		p.AddParkingEntrance("South Gate", "no.2")
		p.AddParkingFloor(1, 50)
		p.AddParkingFloor(2, 50)
		p.GetParkingTicket()
		p.GetParkingTicket()
		p.DisplayParking()
		//p.ParkingPanel()*/

	fmt.Println(v, u)
	p.DisplayParking()
	p.RemoveVehicle(t2.ID)
	p.DisplayParking()

}
