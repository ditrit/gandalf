package payload

type RegisterPayload struct {
	name    string
	content string
}

func (rp RegisterPayload) GetName() string {
	return rp.name
}

func (rp RegisterPayload) GetContent() string {
	return rp.content
}
