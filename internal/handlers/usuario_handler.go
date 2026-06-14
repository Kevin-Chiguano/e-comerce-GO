package handlers

import (
	"net/http"
	"ecommerce-manager/internal/services"

	"github.com/gin-gonic/gin"
)

type UsuarioHandler struct {
	service *services.UsuarioService
}

func NewUsuarioHandler(service *services.UsuarioService) *UsuarioHandler {
	return &UsuarioHandler{service: service}
}

// Por ahora vacío, puedes agregar métodos después
func (h *UsuarioHandler) GetProfile(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Perfil de usuario"})
}