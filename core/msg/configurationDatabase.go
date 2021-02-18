package msg

import "github.com/ditrit/shoset/msg"

// ConfigurationDatabase : gandalf configs
type ConfigurationDatabase struct {
	msg.MessageBase
	Target  string
	Command string
	Context map[string]interface{}
}

// NewConfigurationDatabase : Configuration constructor
// todo : passer une map pour gerer les valeurs optionnelles ?
func NewConfigurationDatabase(target string, command string, payload string) *ConfigurationDatabase {
	s := new(ConfigurationDatabase)
	s.InitMessageBase()

	s.Target = target
	s.Context = make(map[string]interface{})
	s.Command = command
	s.Payload = payload
	return s
}

// GetMsgType accessor
func (cd ConfigurationDatabase) GetMsgType() string { return "configurationDatabase" }

// GetTarget :
func (cd ConfigurationDatabase) GetTarget() string { return cd.Target }

// GetCommand :
func (cd ConfigurationDatabase) GetCommand() string { return cd.Command }

// GetContext :
func (cd ConfigurationDatabase) GetContext() map[string]interface{} { return cd.Context }
