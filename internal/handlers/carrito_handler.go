package handlers

import (
	"net/http"
	"ecommerce-manager/internal/services"

	"github.com/gin-gonic/gin"
)

type CarritoHandler struct {
	service *services.CarritoService
}

func NewCarritoHandler(service *services.CarritoService) *CarritoHandler {
	return &CarritoHandler{service: service}
}

func (h *CarritoHandler) GetCarrito(c *gin.Context) {
	usuarioID := 1 // TODO: Más adelante sacar del JWT

	carrito, err := h.service.GetOrCreateCarrito(usuarioID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, carrito)
}

func (h *CarritoHandler) AddItem(c *gin.Context) {
	usuarioID := 1 // TODO: Más adelante sacar del JWT

	var req struct {
		ProductoID int `json:"producto_id"`
		Cantidad   int `json:"cantidad"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.service.AddToCarrito(usuarioID, req.ProductoID, req.Cantidad)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Producto añadido al carrito"})
}