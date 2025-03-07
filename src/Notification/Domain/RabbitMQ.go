package domain

type IMessagePublisher interface {
	PublishMessage(message string) error
}
