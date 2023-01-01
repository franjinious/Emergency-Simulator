package agent

import (
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	"gitlab.utc.fr/wanhongz/emergency-simulator/agent/doctor"
	"gitlab.utc.fr/wanhongz/emergency-simulator/agent/nurse"
	"gitlab.utc.fr/wanhongz/emergency-simulator/agent/patient"
	"gitlab.utc.fr/wanhongz/emergency-simulator/agent/rooms"
)

type Hospital struct {
	sync.Mutex
	ID                  int // 给病人分配的序号
	NurseCenter         *nurse.Nurse_manager
	ReceptionCenter     *rooms.ReceptionRoom
	EmergencyRoomCenter *rooms.EmergencyRoomManager
	DoctorCenter        *doctor.DoctorManager
	WaitingCenter       *rooms.WaitingRoom
}

func CreateHospital() *Hospital {
	ini()
	h := &Hospital{}
	h.ID = 1
	h.NurseCenter = nurse.GetInstance(3)
	h.ReceptionCenter = rooms.GetInstance(5)
	h.EmergencyRoomCenter = rooms.GetEmergencyRoomManagerInstance(10)
	h.DoctorCenter = doctor.GetDoctorManagerInstance(5)
	h.WaitingCenter = rooms.GetWaitingRoomInstance(h.EmergencyRoomCenter.MsgRequest, h.EmergencyRoomCenter.MsgReponse, h.DoctorCenter.DoctorReuqest, h.DoctorCenter.DoctorResponce)
	return h
}

func (h *Hospital) CreatePatient(w http.ResponseWriter, r *http.Request) {
	h.Lock()
	h.AcceptNewPatient(1, true, "111", 10)
	h.Unlock()
}

func ini(){
	bande := " ｜ ╔═╗┌┬┐┌─┐┬─┐┌─┐┌─┐┌┐┌┌─┐┬ ┬   ╔═╗┬┌┬┐┬ ┬┬  ┌─┐┌┬┐┌─┐┬─┐ | \n ｜ ║╣ │││├┤ ├┬┘│ ┬├┤ ││││  └┬┘───╚═╗│││││ ││  ├─┤ │ │ │├┬┘ | \n ｜ ╚═╝┴ ┴└─┘┴└─└─┘└─┘┘└┘└─┘ ┴    ╚═╝┴┴ ┴└─┘┴─┘┴ ┴ ┴ └─┘┴└─ | \n"

	fmt.Println("  ----------------------------------------------------------")
	fmt.Print(bande)
	fmt.Println("  ----------------------------------------------------------")
}

func (h *Hospital) Start() {
	go h.NurseCenter.Run()
	go h.ReceptionCenter.Run()
	go h.EmergencyRoomCenter.Run()
	go h.DoctorCenter.Run()
	go h.WaitingCenter.Run()
	mux := http.NewServeMux()
	mux.HandleFunc("/new", h.CreatePatient)

	s := &http.Server{
		Addr:           "127.0.0.1:8082",
		Handler:        mux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	time.Sleep(1*time.Second)
	fmt.Println("----------------------------------------------------------")
	log.Println("Everything is normal, the emergency room is working")
	log.Fatal(s.ListenAndServe())
}

func (h *Hospital) AcceptNewPatient(age int, gender bool, symptom string, tolerance int) {
	p := &patient.Patient{
		ID:                    h.ID,
		Age:                   age,
		Gender:                gender,
		Symptom:               symptom,
		Severity:              -1,
		Tolerance:             tolerance,
		TimeForTreat:          -1,
		Status:                patient.Waiting_for_nurse,
		Msg_nurse:             make(chan string, 20),
		Msg_request_nurse:     h.NurseCenter.Get_chan_patient(),
		Msg_request_reception: h.ReceptionCenter.MsgRequest,
		Msg_receive_reception: make(chan string, 20),
		Msg_request_waiting:   h.WaitingCenter.MsgRequest,
	}
	h.ID++
	go p.Run()
}
