package rooms

import "gitlab.utc.fr/wanhongz/emergency-simulator/agent/patient"

type WaitingRoom struct {
	MsgRequest    chan *patient.Patient // 接受病人请求的信道
	QueuePatients []*patient.Patient    // 所有等待病人的切片
}
