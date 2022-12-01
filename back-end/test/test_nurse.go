package test

import (
	"fmt"
	"log"
	"time"

	"gitlab.utc.fr/wanhongz/emergency-simulator/agent/nurse"
	"gitlab.utc.fr/wanhongz/emergency-simulator/agent/patient"
	"gitlab.utc.fr/wanhongz/emergency-simulator/agent/rooms"
)

func Test() {
	a := nurse.GetInstance(3)
	b := rooms.GetInstance(5)
	c := rooms.GetEmergencyRoomManagerInstance(4)
	go p(a, b)
	go b.Run()
	go a.Run()
	go c.Run()
	fmt.Scanln()
}

func p(a *nurse.Nurse_manager, b *rooms.ReceptionRoom) {
	time.Sleep(3 * time.Second)
	p := patient.NewPatient(1, 1, true, "111", 1, 1, 1, a.Get_chan_patient(), b.MsgRequest)
	go p.Run()

	p2 := patient.NewPatient(2, 1, true, "111", 1, 1, 1, a.Get_chan_patient(), b.MsgRequest)
	go p2.Run()
	p3 := patient.NewPatient(3, 1, true, "111", 1, 1, 1, a.Get_chan_patient(), b.MsgRequest)
	go p3.Run()

	p4 := patient.NewPatient(4, 1, true, "111", 1, 1, 1, a.Get_chan_patient(), b.MsgRequest)
	go p4.Run()

	p5 := patient.NewPatient(5, 1, true, "111", 1, 1, 1, a.Get_chan_patient(), b.MsgRequest)
	p6 := patient.NewPatient(6, 1, true, "111", 1, 1, 1, a.Get_chan_patient(), b.MsgRequest)
	go p5.Run()
	go p6.Run()

	p7 := patient.NewPatient(7, 1, true, "111", 1, 1, 1, a.Get_chan_patient(), b.MsgRequest)
	p8 := patient.NewPatient(8, 1, true, "111", 1, 1, 1, a.Get_chan_patient(), b.MsgRequest)
	go p7.Run()
	go p8.Run()
	time.Sleep(40 * time.Second)
}

func Test2() {
	r := rooms.GetEmergencyRoomManagerInstance(10)
	go r.Run()
	n(r)
	fmt.Scanln()
}

func n(rr *rooms.EmergencyRoomManager) {
	c_request := rr.MsgRequest
	c_reponse := rr.MsgReponse
	c_request <- 3
	pp := <-c_reponse
	c_request <- 3
	pp2 := <-c_reponse

	log.Println(pp.ID)

	log.Println(pp2.ID)
}
