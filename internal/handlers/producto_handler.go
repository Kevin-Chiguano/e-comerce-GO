package handlers

import (
	"net/http"
	"strconv"
	"ecommerce-manager/internal/services"

	"github.com/gin-gonic/gin"
)

type ProductoHandler struct {
	service *services.ProductoService
}

func NewProductoHandler(service *services.ProductoService) *ProductoHandler {
	return &ProductoHandler{service: service}
}

func (h *ProductoHandler) GetAll(c *gin.Context) {
	productos, err := h.service.GetAllProductos()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, productos)
}

func (h *ProductoHandler) GetByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	producto, err := h.service.GetProductoByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Producto no encontrado"})
		return
	}
	c.JSON(http.StatusOK, producto)
}