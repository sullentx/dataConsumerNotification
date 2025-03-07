package repositories

import domain "api-consumer/Notification/Domain"

type RabbitMQUseCase struct {
	repo domain.IMessagePublisher
}

func NewRabbitMQUseCase(repo domain.IMessagePublisher) *RabbitMQUseCase {
	return &RabbitMQUseCase{repo: repo}
}

func (uc *RabbitMQUseCase) Execute(message string) error {
	return uc.repo.PublishMessage(message)
}
