package handlers

import (
	"net/http"
	"strconv"
	"ecommerce-manager/internal/services"

	"github.com/gin-gonic/gin"
)

type PedidoHandler struct {
	service *services.PedidoService
}

func NewPedidoHandler(service *services.PedidoService) *PedidoHandler {
	return &PedidoHandler{service: service}
}

func (h *PedidoHandler) CrearPedido(c *gin.Context) {
	usuarioID := c.GetInt("user_id") // Obtenido del JWT

	pedido, err := h.service.CrearPedidoDesdeCarrito(usuarioID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Pedido creado exitosamente",
		"pedido":  pedido,
	})
}

func (h *PedidoHandler) GetMisPedidos(c *gin.Context) {
	usuarioID := c.GetInt("user_id") // Obtenido del JWT

	pedidos, err := h.service.GetMisPedidos(usuarioID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Lista de pedidos",
		"pedidos": pedidos,
	})
}

func (h *PedidoHandler) GetAllPedidos(c *gin.Context) {
	pedidos, err := h.service.GetAllPedidos()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Lista de pedidos",
		"pedidos": pedidos,
	})
}

func (h *PedidoHandler) AprobarPedido(c *gin.Context) {
	pedidoID, err := strconv.Atoi(c.Param("id"))
	if err != nil || pedidoID <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID de pedido inválido"})
		return
	}

	err = h.service.AprobarPedido(pedidoID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Pedido aprobado correctamente"})
}