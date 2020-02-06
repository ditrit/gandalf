package constant

//nolint:golint,stylecheck
const (
	COMMAND_READY                      string = "COMMAND_READY"
	COMMAND_WAIT                       string = "COMMAND_WAIT"
	EVENT_WAIT                         string = "EVENT_WAIT"
	COMMAND_MESSAGE                    string = "COMMAND_MESSAGE"
	COMMAND_MESSAGE_REPLY              string = "COMMAND_MESSAGE_REPLY"
	EVENT_VALIDATION_TOPIC             string = "EVENT_VALIDATION_TOPIC"
	COMMAND_VALIDATION_FUNCTIONS       string = "COMMAND_VALIDATION_FUNCTIONS"
	EVENT_VALIDATION_FUNCTIONS         string = "EVENT_VALIDATION_FUNCTIONS"
	COMMAND_VALIDATION_FUNCTIONS_REPLY string = "COMMAND_VALIDATION_FUNCTIONS_REPLY"
	WORKER_SERVICE_CLASS_CAPTURE       string = "WORKER_SERVICE_CLASS_CAPTURE"
)

type State int

//nolint:golint
const (
	ONGOING State = iota
	SUCCES
	FAIL
)
