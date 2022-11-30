package nurse

import (
	"errors"
	"log"
	"strconv"
	"sync"
	"time"

	"gitlab.utc.fr/wanhongz/emergency-simulator/agent/patient"
)

type nurse struct {
	sync.Mutex
	ID       int                   // 护士的唯一ID
	Usable   bool                  // 当前是否空闲
	p        *patient.Patient      // 当前正在判断情况的病人
	manager  *Nurse_manager        // 管理类
	msg_send chan *nurse           // 给管理类发送消息的信道
	msg_recv chan *patient.Patient // 接受管理类的请求
}

// 构造函数
func NewNurse(id int, m *Nurse_manager) *nurse {
	return &nurse{
		ID:       id,
		Usable:   false,
		p:        nil,
		manager:  m,
		msg_send: m.msg_nurse,
		msg_recv: make(chan *patient.Patient, 10),
	}
}

// 接受新的病人
func (n *nurse) TreatNewPatient(p *patient.Patient) error {
	n.Lock()
	if n.Usable != true {
		n.Unlock()
		return errors.New("Patient is not usable")
	} else {
		n.p = p
		n.Unlock()
		return nil
	}
}

func (n *nurse) SetPatient(patient2 *patient.Patient) {
	n.p = patient2
}

// 设置是否空闲
func (n *nurse) SetUsable(b bool) {
	n.Lock()
	n.Usable = b
	n.Unlock()
}

// 设置病人的状态
func (n *nurse) SetPatientStatus(gravite int, time int) {
	n.Lock()
	n.p.SetSeverity(gravite)
	n.p.SetTimeForTreat(time)
	n.Unlock()
}

// 通知管理器 自己空闲 要求分配任务
func (n *nurse) ticket() {
	n.msg_send <- n
}

// 诊断病人
func (n *nurse) judge(patient2 *patient.Patient) {
	// 判断算法

	// sleep模拟处理时间
	patient2.SetStatus(patient.Being_judged_by_nurse)

	time.Sleep(5 * time.Second)
	n.SetPatientStatus(5, 10)

	patient2.Lock()
	patient2.Msg_nurse <- "ticket"
	patient2.Unlock()
}

func (nur *nurse) treat(n *patient.Patient) {
	nur.SetPatient(n)
	nur.SetUsable(false)
	// 调用nurse的函数 设置patient的状态
	nur.judge(n)
	nur.ticket()
}

func (nur *nurse) Run() {
	log.Println("Nurse " + strconv.FormatInt(int64(nur.ID), 10) + " start")
	for {
		select {
		case n := <-nur.msg_recv:
			go nur.treat(n)
		default:
			time.Sleep(1 * time.Second)
		}
	}
}
