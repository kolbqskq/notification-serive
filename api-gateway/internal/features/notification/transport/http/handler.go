package notification_http

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/kolbqskq/notification-service/api-gateway/internal/core/domain"
	pb "github.com/kolbqskq/notification-service/proto/notification/v1"
)

type NotificationService interface {
	SendNotification(ctx context.Context, n *domain.Notification) (uuid.UUID, error)
}

type NotificationHandler struct {
	router              gin.IRouter
	notificationService NotificationService
	historyClient       pb.HistoryServiceClient
}

type NotificationHandlerDeps struct {
	Router              gin.IRouter
	NotificationService NotificationService
	HistoryClient       pb.HistoryServiceClient
}

func NewNotificationHandler(deps NotificationHandlerDeps) {
	h := &NotificationHandler{
		router:              deps.Router,
		notificationService: deps.NotificationService,
		historyClient:       deps.HistoryClient,
	}

	api := h.router.Group("/api/v1")
	api.POST("/notifications", h.sendNotification)
	api.GET("/history", h.getHistory)
	api.GET("/notifications/:id", h.getStatus)
}
