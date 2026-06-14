package handlers

import (
	"net/http"
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
	usuarioID := 1 // TODO: Más adelante sacar del JWT

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
	usuarioID := 1 // TODO: Más adelante sacar del JWT

	c.JSON(http.StatusOK, gin.H{
		"message": "Lista de pedidos (pendiente de implementar)",
		"usuario_id": usuarioID,
	})
}