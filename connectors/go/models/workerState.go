package models

import "sync"

const (
	Ongoing = iota
	Stopping
)

type WorkerState struct {
	sync.Mutex
	state int
}

func NewWorkerState() *WorkerState {
	ws := new(WorkerState)
	ws.state = Ongoing
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
	ws.Lock()
	ws.state = Ongoing
	defer ws.Unlock()
}

func (ws WorkerState) SetStoppingWorkerState() {
	ws.Lock()
	ws.state = Stopping
	defer ws.Unlock()
}
