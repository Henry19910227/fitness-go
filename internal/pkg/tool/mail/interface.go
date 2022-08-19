package mail

type Tool interface {
	Send(recp string, subject string, body string) error
}
