package doctor

import (
	"log"
	"sync"
	"time"
)

type Doctor struct {
	sync.Mutex                  // 互斥量
	ID          int             // 医生的唯一ID
	Specialized map[string]bool // 医生的专业
	Usable      bool            // 当前是否可用
	Ability     int             // 医生的能力数值 1-5分
}

type DoctorManager struct {
	sync.Mutex
	AllDoctor      []*Doctor
	DoctorReuqest  chan int     // 发过来的int类型 代表病人的严重程度
	DoctorResponce chan *Doctor // 返回可用的医生
}

func (dm *DoctorManager) handlerRequest(value int) {
	for i := 0; i < len(dm.AllDoctor); i++ {
		// 依次判断有没有满足的医生
		if dm.AllDoctor[i].Usable == true && dm.AllDoctor[i].Ability >= value {
			dm.DoctorResponce <- dm.AllDoctor[i]
			return
		}
	}
	dm.DoctorResponce <- nil
	return
}

var (
	instance *DoctorManager
	once     sync.Once
)

func GetDoctorManagerInstance(n int) *DoctorManager {
	once.Do(func() {
		instance = &DoctorManager{
			AllDoctor:      make([]*Doctor, 0),
			DoctorReuqest:  make(chan int, 20),
			DoctorResponce: make(chan *Doctor, 20),
		}
		for i := 1; i < n; i++ {
			d := NewDoctor(i, make(map[string]bool), true, i)
			instance.AllDoctor = append(instance.AllDoctor, d)
		}
	})
	return instance
}

func (dm *DoctorManager) Run() {
	log.Println("DoctorManager start working")
	for {
		select {
		case i := <-dm.DoctorReuqest:
			dm.handlerRequest(i)
		default:
			time.Sleep(1 * time.Second)
		}
	}
}

// 构造函数
func NewDoctor(id int, specialized map[string]bool, usable bool, ability int) *Doctor {
	return &Doctor{
		ID:          id,
		Specialized: specialized,
		Usable:      true,
		Ability:     ability,
	}
}

// 设置可用状态
func (d *Doctor) SetUsable(s bool) {
	d.Lock()
	d.Usable = s
	d.Unlock()
}
