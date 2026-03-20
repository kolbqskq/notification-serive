package notification_http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/kolbqskq/notification-service/api-gateway/internal/core/errs"
	pb "github.com/kolbqskq/notification-service/proto/notification/v1"
)

func (h *NotificationHandler) getHistory(c *gin.Context) {
	userID := c.Query("user_id")

	_, err := uuid.Parse(userID)
	if err != nil {
		c.Error(errs.ErrInvalidUserID)
		return
	}

	res, err := h.historyClient.GetHistory(c.Request.Context(), &pb.GetHistoryRequest{
		UserId: userID,
		Limit:  20,
		Offset: 0,
	})
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, res)
}
