package msg

import "github.com/ditrit/shoset/msg"

// Config : gandalf configs
type Secret struct {
	msg.MessageBase
	Target  string
	Command string
	Context map[string]interface{}
}

// NewConfig : Config constructor
// todo : passer une map pour gerer les valeurs optionnelles ?
func NewSecret(target string, command string, payload string) *Secret {
	s := new(Secret)
	s.InitMessageBase()

	s.Target = target
	s.Context = make(map[string]interface{})
	s.Command = command
	s.Payload = payload
	return s
}

// GetMsgType accessor
func (s Secret) GetMsgType() string { return "secret" }

// GetTarget :
func (s Secret) GetTarget() string { return s.Target }

// GetCommand :
func (s Secret) GetCommand() string { return s.Command }

// GetContext :
func (s Secret) GetContext() map[string]interface{} { return s.Context }
