package agent

import "sync"

type patient_status int8

const (
	Waiting_for_nurse      patient_status = 1
	Waiting_for_register   patient_status = 2
	Waiting_for_treat      patient_status = 3
	Being_treated_now      patient_status = 4
	Being_treated_urgently patient_status = 5
)

type patient struct {
	sync.Mutex
	ID           int            // 病人唯一ID
	Age          int            // 年龄
	Gender       bool           // 性别(0 男 1女)
	Symptom      string         // 症状
	Severity     int            // 症状等级 1-10
	Tolerance    int            // 最大等待时间(超过离开/死亡)
	TimeForTreat int            // 预计需要处理的时间(分钟)
	Status       patient_status // 病人当前的状态和进度
}

// 构造函数
func NewPatient(id int, age int, gender bool, symptom string, severity int, tolerance int, timeForTreate int, status patient_status) *patient {
	return &patient{
		ID:           id,
		Age:          age,
		Gender:       gender,
		Symptom:      symptom,
		Severity:     severity,
		Tolerance:    tolerance,
		TimeForTreat: timeForTreate,
		Status:       status,
	}
}

func (p *patient) SetSeverity(s int) {
	p.Lock()
	p.Severity = s
	p.Unlock()
}

func (p *patient) SetTimeForTreat(s int) {
	p.Lock()
	p.TimeForTreat = s
	p.Unlock()
}

func (p *patient) SetStatus(s patient_status) {
	p.Lock()
	p.Status = s
	p.Unlock()
}
