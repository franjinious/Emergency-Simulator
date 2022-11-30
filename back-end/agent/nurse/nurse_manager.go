package nurse

import (
	"log"
	"strconv"
	"sync"
	"time"

	"gitlab.utc.fr/wanhongz/emergency-simulator/agent/patient"
)

// 管理所有的nurse
type Nurse_manager struct {
	sync.Mutex
	now_id              int
	nurse_number        int
	busy_nurse_number   int                   // 空闲护士数量
	usable_nurse_number int                   // 忙碌护士数量
	nurse_pool_busy     []*nurse              // 忙碌队列
	nurse_pool_usable   []*nurse              // 空闲的护士队列
	msg_nurse           chan *nurse           // 护士任务结束的通知信道
	msg_patient         chan *patient.Patient // 病人请求判断情况的信道

}

func (n *Nurse_manager) Get_chan_patient() chan *patient.Patient {
	return n.msg_patient
}

var (
	instance *Nurse_manager
	once     sync.Once
)

// 单例模式
func GetInstance(n int) *Nurse_manager {
	once.Do(func() {
		instance = &Nurse_manager{
			now_id:              0,
			nurse_number:        n,
			busy_nurse_number:   0,
			usable_nurse_number: n,
			nurse_pool_busy:     make([]*nurse, 0),
			nurse_pool_usable:   make([]*nurse, 0),
			msg_nurse:           make(chan *nurse, 10),
			msg_patient:         make(chan *patient.Patient, 10),
		}
		i := 1
		for i <= n {
			a := NewNurse(instance.now_id, instance)
			instance.nurse_pool_usable = append(instance.nurse_pool_usable, a)
			go a.Run()
			instance.now_id++
			i++
		}
	})
	return instance
}

func (nm *Nurse_manager) handler_nurse_request(n *nurse) {
	log.Println("Nurse center gets a nurse " + strconv.FormatInt(int64(n.ID), 10) + " free request")
	// 处理nurse的请求
	nm.Lock()
	for i := 0; i < len(nm.nurse_pool_busy); i++ {
		// 增加可用队列 减少忙碌队列
		if nm.nurse_pool_busy[i] == n {
			nm.nurse_pool_busy = append(nm.nurse_pool_busy[:i], nm.nurse_pool_busy[i+1:]...)
			nm.nurse_pool_usable = append(nm.nurse_pool_usable, n)
			nm.usable_nurse_number++
			nm.busy_nurse_number--
			n.SetUsable(true)
		}
	}

	nm.Unlock()
}

func (nm *Nurse_manager) handler_patient_request(n *patient.Patient) {
	log.Println("Nurse center gets a new patient " + strconv.FormatInt(int64(n.ID), 10) + " request")
	// 处理patient的请求

	// 申请nurse资源
	for {
		nm.Lock()
		if nm.usable_nurse_number > 0 {
			nm.usable_nurse_number--
			nur := nm.nurse_pool_usable[0]

			nm.nurse_pool_usable = nm.nurse_pool_usable[1:]
			nm.nurse_pool_busy = append(nm.nurse_pool_busy, nur)

			nur.msg_recv <- n

			// 结束
			nm.Unlock()
			break
		}
		nm.Unlock()
		time.Sleep(1 * time.Second)
	}
}

func (nm *Nurse_manager) Run() {
	log.Println("Nurse center starting")
	for {
		select {
		case n := <-nm.msg_nurse:
			go nm.handler_nurse_request(n)
		case m := <-nm.msg_patient:
			go nm.handler_patient_request(m)
		default:
			time.Sleep(1 * time.Second)
		}
	}
}
