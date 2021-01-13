package msg

import "github.com/ditrit/shoset/msg"

// Configuration : gandalf configs
type Configuration struct {
	msg.MessageBase
	Target  string
	Command string
	Context map[string]interface{}
}

// NewConfiguration : Configuration constructor
// todo : passer une map pour gerer les valeurs optionnelles ?
func NewConfiguration(target string, command string, payload string) *Configuration {
	s := new(Configuration)
	s.InitMessageBase()

	s.Target = target
	s.Context = make(map[string]interface{})
	s.Command = command
	s.Payload = payload
	return s
}

// GetMsgType accessor
func (s Configuration) GetMsgType() string { return "configuration" }

// GetTarget :
func (s Configuration) GetTarget() string { return s.Target }

// GetCommand :
func (s Configuration) GetCommand() string { return s.Command }

// GetContext :
func (s Configuration) GetContext() map[string]interface{} { return s.Context }
