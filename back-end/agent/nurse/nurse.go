package nurse

import (
	"errors"
	"gitlab.utc.fr/wanhongz/emergency-simulator/agent/patient"
	"sync"
)

type nurse struct {
	sync.Mutex
	ID     int              // 护士的唯一ID
	Usable bool             // 当前是否空闲
	p      *patient.Patient // 当前正在判断情况的病人
	manager *nurse_manager   // 管理类
	msg_send    chan *nurse  // 给管理类发送消息的信道
}

// 构造函数
func NewNurse(id int,m *nurse_manager) *nurse {
	return &nurse{
		ID:     id,
		Usable: false,
		p:      nil,
		manager: m,
		msg_send: m.msg_recv,
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
func (n *nurse) ticket(){
	n.msg_send <- n
}