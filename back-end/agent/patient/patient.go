package patient

import (
	"log"
	"strconv"
	"sync"
	"time"
)

type Patient struct {
	sync.Mutex
	ID                    int            // 病人唯一ID
	Age                   int            // 年龄
	Gender                bool           // 性别(0 男 1女)
	Symptom               string         // 症状
	Severity              int            // 症状等级 1-10
	Tolerance             int            // 最大等待时间(超过离开/死亡)
	TimeForTreat          int            // 预计需要处理的时间(分钟)
	Status                patient_status // 病人当前的状态和进度
	Msg_nurse             chan string    // 接受护士消息的信道
	Msg_request_nurse     chan *Patient  // 请求护士的信道
	Msg_request_reception chan *Patient  // 请求挂号的信道
	Msg_receive_reception chan string    // 接受挂号的信道
	Msg_request_waiting   chan *Patient  // 请求加入等待室候诊队列的信道
	T                     time.Time      // 创建时间
}

// 构造函数
func NewPatient(id int, age int, gender bool, symptom string, severity int, tolerance int, timeForTreate int, c chan *Patient, d chan *Patient, c_w chan *Patient) *Patient {
	return &Patient{
		ID:                    id,
		Age:                   age,
		Gender:                gender,
		Symptom:               symptom,
		Severity:              severity,
		Tolerance:             tolerance,
		TimeForTreat:          timeForTreate,
		Status:                Waiting_for_nurse,
		Msg_nurse:             make(chan string, 20),
		Msg_request_nurse:     c,
		Msg_request_reception: d,
		Msg_receive_reception: make(chan string, 20),
		Msg_request_waiting:   c_w,
		T:                     time.Now(),
	}
}

func (p *Patient) SetSeverity(s int) {
	p.Lock()
	p.Severity = s
	p.Unlock()
}

func (p *Patient) SetTimeForTreat(s int) {
	p.Lock()
	p.TimeForTreat = s
	p.Unlock()
}

func (p *Patient) SetStatus(s patient_status) {
	p.Lock()
	p.Status = s
	p.Unlock()
}

func (p *Patient) RequestCheckingStatus() {
	p.Lock()
	p.Msg_request_nurse <- p
	p.Unlock()
}

func (p *Patient) Run() {
	p.RequestCheckingStatus()
	for {
		select {
		case n := <-p.Msg_nurse:
			if n == "ticket" {
				p.SetStatus(Waiting_for_register)
				log.Println("Patient " + strconv.FormatInt(int64(p.ID), 10) + " get a status: gravity " + strconv.FormatInt(int64(p.Severity), 10) + ", need time " + strconv.FormatInt(int64(p.TimeForTreat), 10) + ", and go to get a reception")
				p.Lock()
				p.Msg_request_reception <- p
				p.Unlock()
			}
		case m := <-p.Msg_receive_reception:
			if m == "ticket" {
				log.Println("Patient " + strconv.FormatInt(int64(p.ID), 10) + " get reception")
				p.SetStatus(Waiting_for_treat)
				p.Lock()
				p.Msg_request_waiting <- p
				p.Unlock()
			}

		default:
			time.Sleep(1 * time.Second)
		}
	}
}
