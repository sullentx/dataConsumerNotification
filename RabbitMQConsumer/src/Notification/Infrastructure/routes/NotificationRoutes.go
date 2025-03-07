package routes

import (
	"api-consumer/Notification/Infrastructure/controller"

	"github.com/gin-gonic/gin"
)

func SetRoutes(router *gin.Engine, postNotification *controller.PostNotificationHandler, getAllNotification *controller.GetAllNotificationHandler) {
	router.POST("/notifications", postNotification.HandlePost)
	router.GET("/notifications", getAllNotification.HandleGetAll)

}
