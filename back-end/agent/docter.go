package agent

import "sync"

type doctor struct {
	sync.Mutex                  // 互斥量
	ID          int             // 医生的唯一ID
	Specialized map[string]bool // 医生的专业
	Usable      bool            // 当前是否可用
	Ability     int             // 医生的能力数值 1-5分
}

// 构造函数
func NewDoctor(id int, specialized map[string]bool, usable bool, ability int) *doctor {
	return &doctor{
		ID:          id,
		Specialized: specialized,
		Usable:      usable,
		Ability:     ability,
	}
}

// 设置可用状态
func (d *doctor) SetUsable(s bool) {
	d.Lock()
	d.Usable = s
	d.Unlock()
}
