package notification_http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/google/uuid"
	"github.com/kolbqskq/notification-service/api-gateway/internal/core/domain"
	"github.com/kolbqskq/notification-service/api-gateway/internal/core/errs"
)

type SendNotificationRequest struct {
	UserID  string `json:"user_id" binding:"required"`
	Message string `json:"message" binding:"required"`
}

func (h *NotificationHandler) sendNotification(c *gin.Context) {
	var req SendNotificationRequest

	if err := c.ShouldBindBodyWith(&req, binding.JSON); err != nil {
		c.Error(err)
		return
	}
	userID, err := uuid.Parse(req.UserID)
	if err != nil {
		c.Error(errs.ErrInvalidUserID)
		return
	}

	notification, err := domain.NewNotification(userID, req.Message)
	if err != nil {
		c.Error(err)
		return
	}

	id, err := h.notificationService.SendNotification(c.Request.Context(), notification)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusAccepted, gin.H{
		"status": "queued",
		"id":     id,
	})
}
