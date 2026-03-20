package notification_http

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/kolbqskq/notification-service/api-gateway/internal/core/errs"
	pb "github.com/kolbqskq/notification-service/proto/notification/v1"
)

func (h *NotificationHandler) getStatus(c *gin.Context) {
	id := c.Param("id")
	_, err := uuid.Parse(id)
	if err != nil {
		c.Error(errs.ErrInvalidUserID)
		return
	}

	res, err := h.historyClient.GetStatus(c.Request.Context(), &pb.GetStatusRequest{
		NotificationId: id,
	})

	if err != nil {
		c.Error(err)
		return
	}

	response := gin.H{
		"status": res.Status,
	}
	if res.SentAt != nil {
		response["sent_at"] = res.SentAt.AsTime().Format(time.RFC3339)
	}
	if res.ErrorMessage != "" {
		response["error_message"] = res.ErrorMessage
	}

	c.JSON(http.StatusOK, response)
}
