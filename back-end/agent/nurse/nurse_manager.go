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
	PatientWaiting      int // 等待的病人数量
	now_id              int
	nurse_number        int
	nurse_pool_busy     []*nurse              // 忙碌队列
	nurse_pool_usable   []*nurse              // 空闲的护士队列
	msg_nurse           chan *nurse           // 护士任务结束的通知信道
	msg_patient         chan *patient.Patient // 病人请求判断情况的信道

}

func (n *Nurse_manager) GetNurseNumber() int {

	return n.nurse_number
}

func (n *Nurse_manager) GetBusyQueue() []*nurse {
	return n.nurse_pool_busy
}

func (n *Nurse_manager) GetfreeQueue() []*nurse {
	return n.nurse_pool_usable
}

// 添加成功返回true 失败false 最大值为5
func (n *Nurse_manager) Add_patient() {
	flag := false
	for flag == false {
		if n.TryLock() {
			flag = true

			// 添加nurse
			n.nurse_number++
			a := NewNurse(n.now_id, instance)
			n.now_id++
			n.nurse_pool_usable = append(n.nurse_pool_usable, a)
			go a.Run()
			n.Unlock()
		}
	}
}

func (n *Nurse_manager) Reduce_patient() {
	n.Lock()

	if len(n.nurse_pool_usable) != 0 {
		t := n.nurse_pool_usable[0]
		t.Lock()

		n.nurse_number--

		n.nurse_pool_usable = n.nurse_pool_usable[1:]
		close(t.GetChan())
		t.Unlock()
	} else {
		t := n.nurse_pool_busy[0]
		t.Lock()
		aa := t.p
		close(t.GetChan())
		t.Unlock()

		n.nurse_number--
		n.nurse_pool_busy = n.nurse_pool_busy[1:]

		aa.RequestCheckingStatus()
	}

	n.Unlock()
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
			PatientWaiting:      0,
			now_id:              1,
			nurse_number:        n,
			nurse_pool_busy:     make([]*nurse, 0),
			nurse_pool_usable:   make([]*nurse, 0),
			msg_nurse:           make(chan *nurse, 20),
			msg_patient:         make(chan *patient.Patient, 20),
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
			n.SetUsable(true)
		}
	}

	nm.Unlock()
}

func (nm *Nurse_manager) handler_patient_request(n *patient.Patient) {
	nm.Lock()
	nm.PatientWaiting++
	nm.Unlock()
	log.Println("Nurse center gets a new patient " + strconv.FormatInt(int64(n.ID), 10) + " request, now there are totally " + strconv.FormatInt(int64(nm.PatientWaiting), 10) + " waiting")
	// 处理patient的请求
	// 申请nurse资源

	for {
		nm.Lock()
		if len(nm.nurse_pool_usable) > 0 {

			nur := nm.nurse_pool_usable[0]

			nm.nurse_pool_usable = nm.nurse_pool_usable[1:]
			nm.nurse_pool_busy = append(nm.nurse_pool_busy, nur)

			nm.PatientWaiting--
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
