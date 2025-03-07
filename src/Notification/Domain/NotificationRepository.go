package domain

import entities "api-consumer/Notification/Domain/Entities"

type NotificationRepository interface {
	Save(notification entities.Notification) error
	GetMessages() ([]entities.Notification, error)
}
