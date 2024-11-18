package middleware

import (
	"context"
	"github.com/gin-gonic/gin"
	"time"
)

func ContextWIthTimeout() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
		defer cancel()
		c.Set("requestContext", ctx)
		c.Next()
	}
}
func GetRequestContext(c *gin.Context) context.Context {
	if ctx, exists := c.Get("requestContext"); exists {
		if reqCtx, ok := ctx.(context.Context); ok {
			return reqCtx
		}
	}
	return context.Background()
}
