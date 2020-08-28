package mail

type MailPayload struct {
	Sender    string
	Body      string
	Receivers []string
	Username  string
	Password  string
	Host      string
}
