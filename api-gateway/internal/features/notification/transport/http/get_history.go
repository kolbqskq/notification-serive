package notification_http

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/kolbqskq/notification-service/api-gateway/internal/core/errs"
	pb "github.com/kolbqskq/notification-service/proto/notification/v1"
)

type NotificationRecordResponse struct {
	ID            string  `json:"id"`
	UserID        string  `json:"user_id"`
	Type          string  `json:"type"`
	Message       string  `json:"message"`
	SourceService string  `json:"source_service"`
	Status        string  `json:"status"`
	CreatedAt     string  `json:"created_at"`
	SentAt        *string `json:"sent_at,omitempty"`
	ErrorMessage  *string `json:"error_message,omitempty"`
}

type GetHistoryResponse struct {
	Records []NotificationRecordResponse `json:"records"`
	Total   int32                        `json:"total"`
}

// @Summary      История уведомлений
// @Description  Возвращает историю уведомлений пользователя
// @Tags         notifications
// @Produce      json
// @Param        user_id query string true "ID пользователя"
// @Param        limit   query int    false "Количество записей (по умолчанию 10)"
// @Param        offset  query int    false "Смещение"
// @Success      200 {object} GetHistoryResponse'
// @Failure      400 {object} ErrorResponse "Bad request"
// @Failure      500 {object} ErrorResponse "Internal server error"
// @Router       /notifications/history [get]
func (h *NotificationHandler) getHistory(c *gin.Context) {
	userID := c.Query("user_id")
	limitStr := c.DefaultQuery("limit", "10")
	offsetStr := c.DefaultQuery("offset", "0")

	_, err := uuid.Parse(userID)
	if err != nil {
		c.Error(errs.ErrInvalidUserID)
		return
	}
	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		limit = 10
	}
	offset, err := strconv.Atoi(offsetStr)
	if err != nil {
		offset = 0
	}

	res, err := h.historyClient.GetHistory(c.Request.Context(), &pb.GetHistoryRequest{
		UserId: userID,
		Limit:  int32(limit),
		Offset: int32(offset),
	})
	if err != nil {
		c.Error(err)
		return
	}

	var records []NotificationRecordResponse

	for _, v := range res.Records {
		r := NotificationRecordResponse{
			ID:            v.Id,
			UserID:        v.UserId,
			Type:          v.Type,
			Message:       v.Message,
			SourceService: v.SourceService,
			Status:        v.Status.String(),
			CreatedAt:     v.CreatedAt.AsTime().Format(time.RFC3339),
		}
		if v.SentAt != nil {
			t := v.SentAt.AsTime().Format(time.RFC3339)
			r.SentAt = &t
		}
		if v.ErrorMessage != "" {
			r.ErrorMessage = &v.ErrorMessage
		}
		records = append(records, r)
	}

	response := GetHistoryResponse{
		Records: records,
		Total:   res.Total,
	}

	c.JSON(http.StatusOK, response)
}
