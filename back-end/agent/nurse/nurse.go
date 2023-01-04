package nurse

import (
	"errors"
	"log"
	"math/big"
	"strconv"
	"sync"
	"time"

	"crypto/rand"

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
		msg_recv: make(chan *patient.Patient, 20),
	}
}

func (n *nurse) GetChan() chan *patient.Patient {
	return n.msg_recv
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
	// n.Lock()
	n.Lock()
	patient2.SetStatus(patient.Being_judged_by_nurse)
	timee, _ := rand.Int(rand.Reader, big.NewInt(5))
	tt := int(timee.Int64())
	tt += 3
	time.Sleep(time.Duration(tt) * time.Second)

	// 严重程度 随机1-5
	gra, _ := rand.Int(rand.Reader, big.NewInt(4))

	// 时间 随机1-10
	tim, _ := rand.Int(rand.Reader, big.NewInt(6))
	n.Unlock()
	if patient2.Severity == -1 {
		n.SetPatientStatus(int(gra.Int64()+1), 10+int(tim.Int64()))
	} else {
		n.SetPatientStatus(patient2.Severity, 5 + 2 * patient2.Severity + int(tim.Int64()))
	}

	patient2.Lock()
	patient2.Msg_nurse <- "ticket"
	patient2.Unlock()
	// n.Unlock()
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
		case n, ok := <-nur.msg_recv:
			if !ok {
				log.Println("Nurse " + strconv.FormatInt(int64(nur.ID), 10) + " stop")
				return
			}
			nur.treat(n)
		default:
			time.Sleep(1 * time.Second)
		}
	}
}
