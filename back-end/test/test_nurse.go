package test

import (
	"gitlab.utc.fr/wanhongz/emergency-simulator/agent/nurse"
	"gitlab.utc.fr/wanhongz/emergency-simulator/agent/patient"
	"time"
)

func Test(){
	a := nurse.GetInstance(3)
	go p(a)
	a.Run()
}

func p(a *nurse.Nurse_manager){
	time.Sleep(3*time.Second)
	p := patient.NewPatient(1,1,true,"111",1,1,1, a.Get_chan_patient())
	go p.Run()

	p2 := patient.NewPatient(2,1,true,"111",1,1,1, a.Get_chan_patient())
	go p2.Run()

	p3 := patient.NewPatient(3,1,true,"111",1,1,1, a.Get_chan_patient())
	go p3.Run()

	p4 := patient.NewPatient(4,1,true,"111",1,1,1, a.Get_chan_patient())
	go p4.Run()
}