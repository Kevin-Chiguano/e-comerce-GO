package middleware

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

// LoggerMiddleware es un middleware personalizado para logging
func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		// Procesar la solicitud
		c.Next()

		// Log después de la respuesta
		end := time.Now()
		latency := end.Sub(start)

		logMsg := fmt.Sprintf("[GIN] %s | %d | %s | %s %s\n",
			end.Format("2006/01/02 - 15:04:05"),
			c.Writer.Status(),
			latency.String(),
			c.Request.Method,
			c.Request.URL.Path,
		)

		gin.DefaultWriter.Write([]byte(logMsg))
	}
}