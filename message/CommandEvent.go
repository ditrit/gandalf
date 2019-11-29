package message

type Event struct {
	topic string
	uuid  string
	acces string
	info  string
}

func (e Event) new(topic, uuid, acces, info string) {
	e.topic = topic
	e.uuid = uuid
	e.acces = acces
	e.info = info
}

func (e Event) sendWith() {

}

func (e Event) from() {

}
