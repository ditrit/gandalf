package worker

type CommandState struct {
	stateValue string
	state      bool
	payload    string
}

func (c CommandState) GetStateValue() string {
	return c.stateValue
}

func (c CommandState) SetStateValue(statevalue string) {
	c.stateValue = statevalue
}

func (c CommandState) GetState() bool {
	return c.state
}

func (c CommandState) SetState(state bool) {
	c.state = state
}

func (c CommandState) GetStatePayload() string {
	return c.payload
}

func (c CommandState) SetStatePayload(payload string) {
	c.payload = payload
}
