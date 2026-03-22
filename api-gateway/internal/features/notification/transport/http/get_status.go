package notification_http

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/kolbqskq/notification-service/api-gateway/internal/core/errs"
	pb "github.com/kolbqskq/notification-service/proto/notification/v1"
)

type GetStatusResponse struct {
	Status       string  `json:"status"`
	SentAt       *string `json:"sent_at,omitempty"`
	ErrorMessage *string `json:"error_message,omitempty"`
}

// @Summary      Получить статус уведомления
// @Description  Делает запрос в history-service по id
// @Tags         notifications
// @Produce      json
// @Param        id path string true "ID уведомления"
// @Success      200  {object} 	GetStatusResponse
// @Failure      400  {object}  ErrorResponse "Bad request"
// @Failure      500  {object}  ErrorResponse "Internal server error"
// @Router       /notifications/{id} [get]
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

	response := GetStatusResponse{
		Status: res.Status,
	}
	if res.SentAt != nil {
		t := res.SentAt.AsTime().Format(time.RFC3339)
		response.SentAt = &t
	}
	if res.ErrorMessage != "" {
		response.ErrorMessage = &res.ErrorMessage
	}

	c.JSON(http.StatusOK, response)
}
