package controller

import (
	application "api-consumer/Notification/Application"
	entities "api-consumer/Notification/Domain/Entities"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PostNotificationHandler struct {
	PostNotificationUseCase *application.PostNotificationUseCase
}

func NewPostNotificationHandler(postNotificationUseCase *application.PostNotificationUseCase) *PostNotificationHandler {
	return &PostNotificationHandler{PostNotificationUseCase: postNotificationUseCase}
}

func (h *PostNotificationHandler) HandlePost(g *gin.Context) {
	var notification entities.Notification
	if err := g.ShouldBindJSON(&notification); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.PostNotificationUseCase.Execute(notification); err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	g.JSON(http.StatusCreated, gin.H{"message": "Notificación enviada"})
}
