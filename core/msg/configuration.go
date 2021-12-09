package msg

import "github.com/ditrit/shoset/msg"

// Configuration : gandalf configs
type Configuration struct {
	msg.MessageBase
	TargetAddress     string
	TargetLogicalName string
	Command           string
	Context           map[string]interface{}
}

// NewConfiguration : Configuration constructor
// todo : passer une map pour gerer les valeurs optionnelles ?
func NewConfiguration(command string, payload string) *Configuration {
	s := new(Configuration)
	s.InitMessageBase()

	s.Context = make(map[string]interface{})
	s.Command = command
	s.Payload = payload
	return s
}

// GetMsgType accessor
func (c Configuration) GetMsgType() string { return "configuration" }

// GetTarget :
func (c Configuration) GetTargetLogicalName() string { return c.TargetLogicalName }

// GetTarget :
func (c Configuration) GetTargetAddress() string { return c.TargetAddress }

// GetCommand :
func (c Configuration) GetCommand() string { return c.Command }

// GetContext :
func (c Configuration) GetContext() map[string]interface{} { return c.Context }
