package test

import (
	"fmt"
	"gitlab.utc.fr/wanhongz/emergency-simulator/agent"
	"log"
	"time"

	"gitlab.utc.fr/wanhongz/emergency-simulator/agent/rooms"
)

func Test() {
	h := agent.CreateHospital()
	h.Start()
	p(h)
	fmt.Scanln()
}

func p(h *agent.Hospital) {
	time.Sleep(3 * time.Second)

	h.AcceptNewPatient(1,true,"111",10)
	h.AcceptNewPatient(1,true,"111",10)
	h.AcceptNewPatient(1,true,"111",10)
	h.AcceptNewPatient(1,true,"111",10)

	time.Sleep(50 * time.Second)
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
