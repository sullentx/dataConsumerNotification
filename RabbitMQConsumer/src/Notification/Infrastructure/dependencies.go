package infraestructure

import (
	core "api-consumer/Core"
	application "api-consumer/Notification/Application"
	"api-consumer/Notification/Application/repositories"
	services "api-consumer/Notification/Application/service"
	"api-consumer/Notification/Infrastructure/adapters"
	"api-consumer/Notification/Infrastructure/controller"
)

var (
	PostNotificationHandler   *controller.PostNotificationHandler
	GetAllNotificationHandler *controller.GetAllNotificationHandler
)

func Init() {

	core.InitPostgres()
	db := core.GetDB()
	RabbitMQAdapter := adapters.InitRabbitMQ()
	rabbitUseCase := repositories.NewRabbitMQUseCase(RabbitMQAdapter)
	rabbitService := services.NewRabbitMQService(rabbitUseCase)
	notificationRepo := NewPostgresNotificationRepository(db)

	postNotificationUseCase := application.NewPostNotificationUseCase(notificationRepo, rabbitService)
	getAllNotificationUseCase := application.GetAllNotification(notificationRepo)
	GetAllNotificationHandler = controller.NewGetAllNotificationHandler(getAllNotificationUseCase)
	PostNotificationHandler = controller.NewPostNotificationHandler(postNotificationUseCase)

}
