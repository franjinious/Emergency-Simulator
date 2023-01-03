package rooms

import (
	"log"
	"strconv"
	"sync"
	"time"

	"gitlab.utc.fr/wanhongz/emergency-simulator/agent/doctor"
	"gitlab.utc.fr/wanhongz/emergency-simulator/agent/patient"
)

type WaitingRoom struct {
	sync.Mutex
	MsgRequest           chan *patient.Patient // 接受病人请求的信道 在监听到病人请求后 将其加入到等待队列中
	QueuePatients        []*patient.Patient    // 所有等待病人的切片
	EmergencyRoomRequest chan int              // 申请急诊室资源的信道
	EmergencyRoomReponse chan *EmergencyRoom   // 急诊室资源的反馈信道
	DoctorsRequest       chan int              // 申请医生的信道
	DoctorsResponse      chan *doctor.Doctor
}

var (
	instance_wr *WaitingRoom
	once3       sync.Once
)

func GetWaitingRoomInstance(c1 chan int, c2 chan *EmergencyRoom, c3 chan int, c4 chan *doctor.Doctor) *WaitingRoom {
	once3.Do(func() {
		instance_wr = &WaitingRoom{
			MsgRequest:           make(chan *patient.Patient, 20),
			QueuePatients:        make([]*patient.Patient, 0),
			EmergencyRoomRequest: c1,
			EmergencyRoomReponse: c2,
			DoctorsRequest:       c3,
			DoctorsResponse:      c4,
		}
	})
	return instance_wr
}

func (wr *WaitingRoom) handlerPatientWaitingRequest(p *patient.Patient) {
	wr.QueuePatients = append(wr.QueuePatients, p)
}

func (wr *WaitingRoom) Run() {
	log.Println("Waiting Room start working")
	for {
		select {
		case i := <-wr.MsgRequest:
			log.Println("Waiting room get a new patient " + strconv.FormatInt(int64(i.ID), 10) + " join request")
			wr.handlerPatientWaitingRequest(i)
		default:
			wr.work()
		}
	}
}

// 遍历病人 寻找资源足够的病人进行治疗 高等级优先
// 对于某一个病人
// 请求急诊室管理类 申请急诊室资源
// 请求医生管理类 申请医生资源
// 可被处理的病人按优先级顺序执行
func (wr *WaitingRoom) work() {
	for i := 0; i < len(wr.QueuePatients); i++ {
		wr.DoctorsRequest <- wr.QueuePatients[i].Severity
		p1 := <-wr.DoctorsResponse
		wr.EmergencyRoomRequest <- wr.QueuePatients[i].Severity
		p2 := <-wr.EmergencyRoomReponse

		// 如果两个资源同时都是可以获取的 则进行治疗
		if p1 != nil && p2 != nil {
			pp := wr.QueuePatients[i]
			p1.Lock()
			p1.Usable = false
			p1.Unlock()
			p2.Lock()
			p2.Status = 1
			p2.Unlock()
			pp.Status = patient.Being_treated_now
			wr.QueuePatients = append(wr.QueuePatients[:i], wr.QueuePatients[i+1:]...)
			go treat(p2, p1, pp)
		}
	}
}

func treat(r *EmergencyRoom, d *doctor.Doctor, p *patient.Patient) {
	// 模拟治疗时间
	log.Println("patient " + strconv.FormatInt(int64(p.ID), 10) + " start to be treat by doctor " + strconv.FormatInt(int64(d.ID), 10) + " in room " + strconv.FormatInt(int64(r.ID), 10))
	time.Sleep(10 * time.Second)
	p.Status = patient.Finish
	log.Println("patient " + strconv.FormatInt(int64(p.ID), 10) + " finish treat ")

	// Alltime["patient " + strconv.FormatInt(int64(p.ID), 10)] = p.T.Sub(time.Now()).Seconds()

	r.Lock()
	r.Status = 0
	r.Unlock()
	d.Lock()
	d.Usable = true
	d.Unlock()
}
