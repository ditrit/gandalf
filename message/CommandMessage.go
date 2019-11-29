package message

type Command struct {
	uuid    string
	routing string
	acces   string
	info    string
}

func (c Command) new(uuid, routing, acces, info string) {
	c.uuid = uuid
	c.routing = routing
	c.acces = acces
	c.info = info
}

func (c Command) sendWith() {

}

func (c Command) from() {

}

func (c Command) response() {

}

type CommandResponse struct {
	command Command
}

func (cr CommandResponse) new(uuid, routing, acces, info string) {
	cr.command.uuid = uuid
	cr.command.routing = routing
	cr.command.acces = acces
	cr.command.info = info
}

func (cr CommandResponse) sendWith() {
}

func (cr CommandResponse) from() {

}
