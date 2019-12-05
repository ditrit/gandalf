package worker

type CommandState struct {
	stateValue string
	state      bool
	payload    string
}

func (c CommandState) GetStateValue() (stateValue string, err error) {
	return c.stateValue
}

func (c CommandState) SetStateValue(statevalue string) err error {
	c.stateValue = statevalue
}

func (c CommandState) GetState() (state bool, err error) {
	return c.state
}

func (c CommandState) SetState(state bool) err error {
	c.state = state
}

func (c CommandState) GetStatePayload() (payload string, err error) {
	return c.payload
}

func (c CommandState) SetStatePayload(payload string) err error {
	c.payload = payload
}
