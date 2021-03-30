package msg

import "github.com/ditrit/shoset/msg"

// Configuration : gandalf configs
type LogicalConfiguration struct {
	msg.MessageBase
	Target  string
	Command string
	Context map[string]interface{}
}

// NewConfiguration : Configuration constructor
// todo : passer une map pour gerer les valeurs optionnelles ?
func NewLogicalConfiguration(target string, command string, payload string) *Configuration {
	s := new(Configuration)
	s.InitMessageBase()

	s.Target = target
	s.Context = make(map[string]interface{})
	s.Command = command
	s.Payload = payload
	return s
}

// GetMsgType accessor
func (lc LogicalConfiguration) GetMsgType() string { return "logicalConfiguration" }

// GetTarget :
func (lc LogicalConfiguration) GetTarget() string { return lc.Target }

// GetCommand :
func (lc LogicalConfiguration) GetCommand() string { return lc.Command }

// GetContext :
func (lc LogicalConfiguration) GetContext() map[string]interface{} { return lc.Context }
