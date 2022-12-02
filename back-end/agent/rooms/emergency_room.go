package rooms

import (
	"log"
	"strconv"
	"sync"
	"time"
)

// 普通会诊室

type EmergencyRoom struct {
	sync.Mutex
	Level  int // 病房的等级 病人只能使用等比自己低的病房
	Status int // 0空闲 1占用
	ID     int // 唯一ID
}

// waitingRoom 在遍历过程中 确定病人的需求后
// 向急诊室管理agent发送请求 判断有无对应类型的房间
type EmergencyRoomManager struct {
	sync.Mutex
	RoomList   map[string]*EmergencyRoom // 所有的房间
	MsgRequest chan int                  // 用于接收房间请求的信道 int 代表所需房间的等级
	MsgReponse chan *EmergencyRoom       // 响应信道 用于返回可用的房间没有返回nil
}

func (erm *EmergencyRoomManager) Run() {
	log.Println("EmergencyRoomManager start working")
	for {
		select {
		case i := <-erm.MsgRequest:
			log.Println("EmergencyRoomCenter get a new request of level " + strconv.FormatInt(int64(i), 10))
			// 不能用协程
			erm.check(i)
		default:
			time.Sleep(1 * time.Second)
		}
	}
}

var (
	instance_erm *EmergencyRoomManager
	once2        sync.Once
)

func GetEmergencyRoomManagerInstance(n int) *EmergencyRoomManager {
	once2.Do(func() {
		instance_erm = &EmergencyRoomManager{
			RoomList:   make(map[string]*EmergencyRoom),
			MsgRequest: make(chan int, 10),
			MsgReponse: make(chan *EmergencyRoom, 10),
		}

		for i := 1; i <= n; i++ {
			instance_erm.RoomList["EmergencyRoom"+strconv.FormatInt(int64(i), 10)] = &EmergencyRoom{
				Level:  i,
				Status: 0,
				ID:     i,
			}
		}
	})
	return instance_erm
}

func (erm *EmergencyRoomManager) check(LevelNeed int) {
	erm.Lock()
	var ans *EmergencyRoom = nil
	for i := 1; i <= LevelNeed; i++ {
		// 判断是否空闲
		r := erm.RoomList["EmergencyRoom"+strconv.FormatInt(int64(i), 10)]
		r.Lock()
		if r.Status == 0 {
			ans = r
		}
		r.Unlock()
	}
	erm.MsgReponse <- ans
	erm.Unlock()
}
