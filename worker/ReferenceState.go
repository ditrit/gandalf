package worker

type ReferenceState struct {
	state Constant.State
}

func (r ReferenceState) GetState() Constant.State {
	return r.state
}

func (r ReferenceState) New(state Constant.State) ReferenceState {
	r.state = state
}
