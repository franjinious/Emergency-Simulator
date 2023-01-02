package agent

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"
	"time"

	"gitlab.utc.fr/wanhongz/emergency-simulator/agent/doctor"
	"gitlab.utc.fr/wanhongz/emergency-simulator/agent/nurse"
	"gitlab.utc.fr/wanhongz/emergency-simulator/agent/patient"
	"gitlab.utc.fr/wanhongz/emergency-simulator/agent/rooms"
)

type Hospital struct {
	sync.Mutex
	IP                  string
	Port                string
	ID                  int // 给病人分配的序号
	NurseCenter         *nurse.Nurse_manager
	ReceptionCenter     *rooms.ReceptionRoom
	EmergencyRoomCenter *rooms.EmergencyRoomManager
	DoctorCenter        *doctor.DoctorManager
	WaitingCenter       *rooms.WaitingRoom
}

func CreateHospital(i string, p string) *Hospital {
	ini()
	h := &Hospital{}
	h.IP = i
	h.Port = p
	h.ID = 1
	h.NurseCenter = nurse.GetInstance(3)
	h.ReceptionCenter = rooms.GetInstance(5)
	h.EmergencyRoomCenter = rooms.GetEmergencyRoomManagerInstance(10)
	h.DoctorCenter = doctor.GetDoctorManagerInstance(5)
	h.WaitingCenter = rooms.GetWaitingRoomInstance(h.EmergencyRoomCenter.MsgRequest, h.EmergencyRoomCenter.MsgReponse, h.DoctorCenter.DoctorReuqest, h.DoctorCenter.DoctorResponce)
	return h
}

type requestBody struct {
	Test int `json:"test"`
}

func (h *Hospital) CreatePatient(w http.ResponseWriter, r *http.Request) {
	h.Lock()
	q := r.URL.Query()
	re := q.Get("test")
	c, _ := strconv.Atoi(re)
	h.AcceptNewPatient(1, true, "111", 10, c)
	w.WriteHeader(http.StatusOK)
	// fmt.Println(c)
	h.Unlock()
}

func ini() {
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
	mux.HandleFunc("/createPatient", h.CreatePatient)
	mux.HandleFunc("/activerAccueil", h.ActiverAccueil)
	mux.HandleFunc("/desactiverAccueil", h.DesactiverAccueil)
	mux.HandleFunc("/activerInfirmier", h.ActiverInfirmier)
	mux.HandleFunc("/desactiverInfirmier", h.DesactiverInfirmier)
	mux.HandleFunc("/activerSalle", h.ActiverSalle)
	mux.HandleFunc("/desactiverSalle", h.DesactiverSalle)
	mux.HandleFunc("/activerDoc", h.ActiverDoc)
	mux.HandleFunc("/desactiverDoc", h.DesactiverDoc)
	mux.HandleFunc("/getinfo", h.Getinfo)

	s := &http.Server{
		Addr:           h.IP + ":" + h.Port,
		Handler:        mux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	time.Sleep(1 * time.Second)
	fmt.Println("----------------------------------------------------------")
	log.Println("Everything is normal, the emergency room is working")
	log.Fatal(s.ListenAndServe())
}

func (h *Hospital) AcceptNewPatient(age int, gender bool, symptom string, tolerance int, gra int) {
	p := &patient.Patient{
		ID:                    h.ID,
		Age:                   age,
		Gender:                gender,
		Symptom:               symptom,
		Severity:              gra,
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

func (h *Hospital) ActiverAccueil(writer http.ResponseWriter, request *http.Request) {
	h.Lock()
	h.NurseCenter.Add_patient()
	h.Unlock()
}

func (h *Hospital) DesactiverAccueil(writer http.ResponseWriter, request *http.Request) {
	h.Lock()
	h.NurseCenter.Reduce_patient()
	h.Unlock()
}

func (h *Hospital) DesactiverInfirmier(writer http.ResponseWriter, request *http.Request) {
	h.Lock()
	h.ReceptionCenter.ReduceQueue()
	h.Unlock()
}

func (h *Hospital) ActiverSalle(writer http.ResponseWriter, request *http.Request) {
	h.Lock()
	// h.DoctorCenter.DeleteDoctor()
	q := request.URL.Query()
	re := q.Get("test")
	c, _ := strconv.Atoi(re)
	h.EmergencyRoomCenter.AddRoom(c)
	h.Unlock()
}

func (h *Hospital) ActiverInfirmier(writer http.ResponseWriter, request *http.Request) {
	h.Lock()
	h.ReceptionCenter.AddQueue()
	h.Unlock()
}

func (h *Hospital) DesactiverSalle(writer http.ResponseWriter, request *http.Request) {
	h.Lock()
	h.EmergencyRoomCenter.DeleteRoom()
	h.Unlock()
}

func (h *Hospital) ActiverDoc(writer http.ResponseWriter, request *http.Request) {
	h.Lock()
	// h.DoctorCenter.DeleteDoctor()
	q := request.URL.Query()
	re := q.Get("test")
	c, _ := strconv.Atoi(re)
	h.DoctorCenter.AddDoctor(c)
	h.Unlock()
}

func (h *Hospital) DesactiverDoc(writer http.ResponseWriter, request *http.Request) {
	h.Lock()
	// h.DoctorCenter.DeleteDoctor()
	q := request.URL.Query()
	re := q.Get("test")
	c, _ := strconv.Atoi(re)
	h.DoctorCenter.DeleteDoctor(c)
	h.Unlock()
}

func (h *Hospital) Getinfo(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Access-Control-Allow-Origin", "*")

	h.Lock()

	var re [][]int

	// Accueil
	h.NurseCenter.Lock()

	a := make([]int, len(h.NurseCenter.GetBusyQueue())+len(h.NurseCenter.GetfreeQueue()))

	// fmt.Println(h.NurseCenter.GetBusyQueue())
	for i := 1; i <= len(h.NurseCenter.GetBusyQueue()); i++ {
		a[i-1] = 1
	}

	h.NurseCenter.Unlock()

	// Infirmier
	h.ReceptionCenter.Lock()

	b := make([]int, h.ReceptionCenter.QueueNumber)

	for i := 1; i < h.ReceptionCenter.QueueNumber; i++ {
		b[i-1] = h.ReceptionCenter.QueuesDoctor["Queue"+strconv.FormatInt(int64(i), 10)].Getstatus()
	}

	h.ReceptionCenter.Unlock()

	// patient

	h.NurseCenter.Lock()

	c := make([]int, 1)
	c[0] = len(h.WaitingCenter.QueuePatients)
	h.NurseCenter.Unlock()

	// room
	h.EmergencyRoomCenter.Lock()

	d := make([]int, h.EmergencyRoomCenter.WorkNumber)
	for i := 1; i <= h.EmergencyRoomCenter.WorkNumber; i++ {
		d[i-1] = h.EmergencyRoomCenter.RoomList["EmergencyRoom"+strconv.FormatInt(int64(i), 10)].Status
	}

	h.EmergencyRoomCenter.Unlock()

	h.DoctorCenter.Lock()

	e := make([]int, 1)
	e[0] = len(h.DoctorCenter.AllDoctor)

	h.DoctorCenter.Unlock()

	h.Unlock()

	re = append(re, a)
	re = append(re, b)
	re = append(re, c)
	re = append(re, d)
	re = append(re, e)

	fmt.Println(re)

	serial, _ := json.Marshal(re)
	writer.WriteHeader(http.StatusOK)
	writer.Write(serial)

	return
}
