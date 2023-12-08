package log

import (
	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		_ = ctx.Writer.Status()
	}
}
