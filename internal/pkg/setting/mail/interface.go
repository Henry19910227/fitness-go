package mail

type Setting interface {
	SMTPHostName() string
	Port() string
	Sender() string
	Password() string
}
