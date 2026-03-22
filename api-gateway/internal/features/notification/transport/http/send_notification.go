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

type SendNotificationResponse struct {
	NotificationID string `json:"id"`
	Status         string `json:"status"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

// @Summary      Отправить уведомление
// @Description  Принимает запрос и публикует событие в Kafka
// @Tags         notifications
// @Accept       json
// @Produce      json
// @Param        request body SendNotificationRequest true "SendNotification тело запроса"
// @Success      202  {object} SendNotificationResponse
// @Failure      400  {object}  ErrorResponse "Bad request"
// @Failure      429  {object}  ErrorResponse "Too many requests"
// @Failure      500  {object}  ErrorResponse "Internal server error"
// @Router       /notifications [post]
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

	c.JSON(http.StatusAccepted, SendNotificationResponse{
		NotificationID: id.String(),
		Status:         "queued",
	})
}
