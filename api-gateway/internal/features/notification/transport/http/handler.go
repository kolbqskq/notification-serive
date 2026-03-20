package notification_http

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/kolbqskq/notification-service/api-gateway/internal/core/domain"
)

type NotificationService interface {
	SendNotification(ctx context.Context, n *domain.Notification) error
}

type NotificationHandler struct {
	router              gin.IRouter
	notificationService NotificationService
}

type NotificationHandlerDeps struct {
	Router              gin.IRouter
	NotificationService NotificationService
}

func NewNotificationHandler(deps NotificationHandlerDeps) {
	h := &NotificationHandler{
		router:              deps.Router,
		notificationService: deps.NotificationService,
	}

	api := h.router.Group("/api/v1")
	api.POST("/notifications", h.sendNotification)
}
