package payload

type ExecutePayload struct {
	name string
}

func (ep ExecutePayload) GetName() string {
	return ep.name
}
