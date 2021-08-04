package msg

import "github.com/ditrit/shoset/msg"

// Configuration : gandalf configs
type LogicalConfiguration struct {
	msg.MessageBase
	TargetAddress     string
	TargetLogicalName string
	Command           string
	Context           map[string]interface{}
}

// NewConfiguration : Configuration constructor
// todo : passer une map pour gerer les valeurs optionnelles ?
func NewLogicalConfiguration(command string, payload string) *LogicalConfiguration {
	s := new(LogicalConfiguration)
	s.InitMessageBase()

	s.Context = make(map[string]interface{})
	s.Command = command
	s.Payload = payload
	return s
}

// GetMsgType accessor
func (lc LogicalConfiguration) GetMsgType() string { return "logicalConfiguration" }

// GetTarget :
func (lc LogicalConfiguration) GetTargetLogicalName() string { return lc.TargetLogicalName }

// GetTarget :
func (lc LogicalConfiguration) GetTargetAddress() string { return lc.TargetAddress }

// GetCommand :
func (lc LogicalConfiguration) GetCommand() string { return lc.Command }

// GetContext :
func (lc LogicalConfiguration) GetContext() map[string]interface{} { return lc.Context }
