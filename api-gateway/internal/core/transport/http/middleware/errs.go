package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/kolbqskq/notification-service/api-gateway/internal/core/errs"
	notification_http "github.com/kolbqskq/notification-service/api-gateway/internal/features/notification/transport/http"
	"github.com/rs/zerolog"
)

func ErrorMiddleware(logger zerolog.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if len(c.Errors) == 0 {
			return
		}
		err := c.Errors.Last().Err
		logger.Error().Err(err).Str("method", c.Request.Method).Str("path", c.Request.URL.Path).Msg("request error")

		httpErr := errs.ToHTTPError(err)
		c.JSON(httpErr.Code, notification_http.ErrorResponse{
			Error: httpErr.Message,
		})
	}
}
