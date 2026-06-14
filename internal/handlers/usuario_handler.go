package handlers

import (
	"net/http"
	"ecommerce-manager/internal/models"
	"ecommerce-manager/internal/services"

	"github.com/gin-gonic/gin"
)

type UsuarioHandler struct {
	service *services.UsuarioService
	authService *services.AuthService
}

func NewUsuarioHandler(service *services.UsuarioService, authService *services.AuthService) *UsuarioHandler {
	return &UsuarioHandler{
		service:     service,
		authService: authService,
	}
}

func (h *UsuarioHandler) Register(c *gin.Context) {
	var user models.Usuario
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if user.RolID == 0 {
		user.RolID = 2 // CLIENTE por defecto
	}

	err := h.authService.Register(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Usuario registrado correctamente"})
}