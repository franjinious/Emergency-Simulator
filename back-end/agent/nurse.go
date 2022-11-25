package agent

import (
	"errors"
	"sync"
)

type nurse struct {
	sync.Mutex
	ID     int      // 护士的唯一ID
	Usable bool     // 当前是否空闲
	p      *patient // 当前正在判断情况的病人
}

// 构造函数
func NewNurse(id int) *nurse {
	return &nurse{
		ID:     id,
		Usable: false,
		p:      nil,
	}
}

// 接受新的病人
func (n *nurse) TreatNewPatient(p *patient) error {
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
	n.Usable = b
}

// 设置病人的状态
func (n *nurse) SetPatientStatus(gravite int, time int) {
	n.p.SetSeverity(gravite)
	n.p.SetTimeForTreat(time)
}
