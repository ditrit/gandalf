package models

import (
	"sync"
)

const (
	ongoing = iota
	stopping
)

type WorkerState struct {
	sync.Mutex
	state int
}

func NewWorkerState() *WorkerState {
	ws := new(WorkerState)
	ws.state = ongoing
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

func (ws *WorkerState) SetOngoingWorkerState() {
	ws.Lock()
	ws.state = ongoing
	defer ws.Unlock()
}

func (ws *WorkerState) SetStoppingWorkerState() {
	ws.Lock()
	ws.state = stopping
	defer ws.Unlock()
}
