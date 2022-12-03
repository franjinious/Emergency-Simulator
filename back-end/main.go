package main

import (
	"gitlab.utc.fr/wanhongz/emergency-simulator/agent"
)

func main() {
	var h *agent.Hospital = agent.CreateHospital()
	h.Start()
}
