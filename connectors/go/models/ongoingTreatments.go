package models

import "sync"

type OngoingTreatments struct {
	sync.Mutex
	index int
}

func NewOngoingTreatments() *OngoingTreatments {
	og := new(OngoingTreatments)
	og.index = 0
	return og
}

func (ot *OngoingTreatments) GetIndex() int {
	ot.Lock()
	index := ot.index
	defer ot.Unlock()
	return index
}

/* func (ot OngoingTreatments) setOngoingTreatments(index int) {
	ot.Lock()
	ot.index = index
	defer ot.Unlock()
} */

func (ot *OngoingTreatments) IncrementOngoingTreatments() {
	ot.Lock()
	ot.index++
	defer ot.Unlock()
}

func (ot *OngoingTreatments) DecrementOngoingTreatments() {
	ot.Lock()
	ot.index--
	defer ot.Unlock()
}
