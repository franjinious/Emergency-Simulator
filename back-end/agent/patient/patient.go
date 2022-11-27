package patient

import "sync"

type Patient struct {
	sync.Mutex
	ID           int            // 病人唯一ID
	Age          int            // 年龄
	Gender       bool           // 性别(0 男 1女)
	Symptom      string         // 症状
	Severity     int            // 症状等级 1-10
	Tolerance    int            // 最大等待时间(超过离开/死亡)
	TimeForTreat int            // 预计需要处理的时间(分钟)
	Status       patient_status // 病人当前的状态和进度
	Msg_nurse    chan string    // 接受护士消息的信道
}

// 构造函数
func NewPatient(id int, age int, gender bool, symptom string, severity int, tolerance int, timeForTreate int) *Patient {
	return &Patient{
		ID:           id,
		Age:          age,
		Gender:       gender,
		Symptom:      symptom,
		Severity:     severity,
		Tolerance:    tolerance,
		TimeForTreat: timeForTreate,
		Status:       Waiting_for_nurse,
		Msg_nurse:    make(chan string,10),
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
