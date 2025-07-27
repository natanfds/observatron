package interfaces

type Webhook interface {
	Send(message string) error
}
