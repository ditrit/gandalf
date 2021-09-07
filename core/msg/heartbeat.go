package msg

import "github.com/ditrit/shoset/msg"

// Heartbeat : gandalf heartbeat
type Heartbeat struct {
	msg.MessageBase
	Event   string
	Context map[string]interface{}
}

// NewHeartbeat : Heartbeat constructor
// todo : passer une map pour gerer les valeurs optionnelles ?
func NewHeartbeat(event string) *Heartbeat {
	s := new(Heartbeat)
	s.InitMessageBase()

	s.Context = make(map[string]interface{})
	s.Event = event
	return s
}

// GetMsgType accessor
func (h Heartbeat) GetMsgType() string { return "heartbeat" }

// GetEvent :
func (h Heartbeat) GetEvent() string { return h.Event }

// GetContext :
func (h Heartbeat) GetContext() map[string]interface{} { return h.Context }
