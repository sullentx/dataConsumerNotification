package application

import (
	services "api-consumer/Notification/Application/service"
	domain "api-consumer/Notification/Domain"
	entities "api-consumer/Notification/Domain/Entities"
)

type PostNotificationUseCase struct {
	notificationRepo domain.NotificationRepository
	rabbitService    *services.RabbitMQService
}

func NewPostNotificationUseCase(notificationRepo domain.NotificationRepository,
	rabbitService *services.RabbitMQService) *PostNotificationUseCase {
	return &PostNotificationUseCase{
		notificationRepo: notificationRepo,
		rabbitService:    rabbitService}
}

func (uc *PostNotificationUseCase) Execute(notification entities.Notification) error {
	uc.rabbitService.SendMessage("Notificacion recibida y guardada con exito")
	return uc.notificationRepo.Save(notification)
}
