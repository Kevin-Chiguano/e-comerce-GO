package handlers

import (
	"fmt"
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
	usuarioID := c.GetInt("user_id")

	carrito, err := h.service.GetOrCreateCarrito(usuarioID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, carrito)
}

func (h *CarritoHandler) AddItem(c *gin.Context) {
	usuarioID := c.GetInt("user_id")

	var req struct {
		ProductoID int `json:"producto_id"`
		Cantidad   int `json:"cantidad"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if req.ProductoID <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID de producto inválido"})
		return
	}

	if req.Cantidad <= 0 {
		req.Cantidad = 1
	}

	err := h.service.AddToCarrito(usuarioID, req.ProductoID, req.Cantidad)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Producto añadido correctamente"})
}

func (h *CarritoHandler) RemoveItem(c *gin.Context) {
	usuarioID := c.GetInt("user_id")
	productoIDStr := c.Param("producto_id")

	var productoID int
	if _, err := fmt.Sscanf(productoIDStr, "%d", &productoID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID de producto inválido"})
		return
	}

	err := h.service.RemoveFromCarrito(usuarioID, productoID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Producto removido del carrito"})
}

func (h *CarritoHandler) ClearCarrito(c *gin.Context) {
	usuarioID := c.GetInt("user_id")

	err := h.service.ClearCarrito(usuarioID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Carrito vaciado correctamente"})
}