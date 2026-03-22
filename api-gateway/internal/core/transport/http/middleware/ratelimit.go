package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/kolbqskq/notification-service/api-gateway/internal/core/errs"
	"github.com/redis/go-redis/v9"
)

func RateLimit(rdb *redis.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		key := "rate:" + c.ClientIP()
		ctx := c.Request.Context()
		count, err := rdb.Incr(ctx, key).Result()
		if err != nil {
			c.Error(err)
			c.Abort()
			return
		}

		if count == 1 {
			rdb.Expire(ctx, key, time.Minute)
		}

		if count > 5 {
			c.Error(errs.ErrTooManyRequests)
			c.Abort()
			return
		}

		c.Next()
	}
}
