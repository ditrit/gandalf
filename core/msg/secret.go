package msg

import "github.com/ditrit/shoset/msg"

// Config : gandalf configs
type Secret struct {
	msg.MessageBase
	TargetAddress     string
	TargetLogicalName string
	Command           string
	Context           map[string]interface{}
}

// NewConfig : Config constructor
// todo : passer une map pour gerer les valeurs optionnelles ?
func NewSecret(command string, payload string) *Secret {
	s := new(Secret)
	s.InitMessageBase()

	s.Context = make(map[string]interface{})
	s.Command = command
	s.Payload = payload
	return s
}

// GetMsgType accessor
func (s Secret) GetMsgType() string { return "secret" }

// GetTarget :
func (s Secret) GetTargetLogicalName() string { return s.TargetLogicalName }

// GetTarget :
func (s Secret) GetTargetAddress() string { return s.TargetAddress }

// GetCommand :
func (s Secret) GetCommand() string { return s.Command }

// GetContext :
func (s Secret) GetContext() map[string]interface{} { return s.Context }
