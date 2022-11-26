package nurse

import (
	"sync"
)

// 管理所有的nurse
type nurse_manager struct {
	sync.Mutex
	now_id int
	nurse_number int
	busy_nurse_number int // 空闲护士数量
	usable_nurse_number int // 忙碌护士数量
	nurse_pool_busy []*nurse // 忙碌队列
	nurse_pool_usable []*nurse // 空闲的护士队列
	msg_recv chan *nurse
}

var (
	instance *nurse_manager
	once sync.Once
)

// 单例模式
func GetInstance(n int) *nurse_manager {
	once.Do(func() {
		instance = &nurse_manager{
			now_id: 0,
			nurse_number: n,
			busy_nurse_number: 0,
			usable_nurse_number: n,
			nurse_pool_busy: make([]*nurse,0),
			nurse_pool_usable: make([]*nurse,0),
			msg_recv: make(chan *nurse,10),
		}
		i := 1
		for i <= n {
			instance.nurse_pool_usable = append(instance.nurse_pool_usable, NewNurse(instance.now_id,instance))
			instance.now_id++
		}
	})
	return instance
}

func (nm *nurse_manager) handler(n *nurse){
	// 处理nurse的请求
}

func (nm *nurse_manager) run(){
	select {
	case n := <-nm.msg_recv:
		nm.handler(n)
	default:
	}
}