package middleware

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

// RoleMiddleware valida que el usuario tenga uno de los roles permitidos
func RoleMiddleware(allowedRoles ...int) gin.HandlerFunc {
	return func(c *gin.Context) {
		rolID, exists := c.Get("rol_id")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Rol no encontrado"})
			c.Abort()
			return
		}

		userRole := rolID.(int)
		hasAccess := false

		for _, allowedRole := range allowedRoles {
			if userRole == allowedRole {
				hasAccess = true
				break
			}
		}

		if !hasAccess {
			c.JSON(http.StatusForbidden, gin.H{"error": "No tienes permiso para acceder a este recurso"})
			c.Abort()
			return
		}

		c.Next()
	}
}

// Constantes para roles
const (
	RoleAdmin = 1
	RoleUser  = 2
)
