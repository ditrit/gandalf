package models

import "sync"

type WorkerState struct {
	sync.Mutex
	state int
}

func NewWorkerState() *WorkerState {
	ws := new(WorkerState)
	ws.state = 0
	return ws
}

func (ws WorkerState) GetState() int {
	ws.Lock()
	workerState := ws.state
	defer ws.Unlock()
	return workerState
}

/*
func (ws WorkerState) setWorkerState(state int) {
	ws.Lock()
	ws.state = state
	defer ws.Unlock()
} */

func (ws WorkerState) SetOngoingWorkerState() {
	//ws.Lock()
	ws.state = 0
	//defer ws.Unlock()
}

func (ws WorkerState) SetStoppingWorkerState() {
	//ws.Lock()
	ws.state = 1
	//defer ws.Unlock()
}
