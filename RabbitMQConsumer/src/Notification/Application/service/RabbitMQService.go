package services

import "api-consumer/Notification/Application/repositories"

type RabbitMQService struct {
	rabbitMQUseCase *repositories.RabbitMQUseCase
}

func NewRabbitMQService(rabbitMQUseCase *repositories.RabbitMQUseCase) *RabbitMQService {
	return &RabbitMQService{rabbitMQUseCase: rabbitMQUseCase}
}

func (svc *RabbitMQService) SendMessage(message string) error {
	return svc.rabbitMQUseCase.Execute(message)
}
