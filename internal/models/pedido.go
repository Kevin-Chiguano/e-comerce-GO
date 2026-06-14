package models

import "time"

type Pedido struct {
	ID         int             `json:"id"`
	UsuarioID  int             `json:"usuario_id"`
	Total      float64         `json:"total"`
	Estado     string          `json:"estado"`
	Fecha      time.Time       `json:"fecha"`
	Detalles   []DetallePedido `json:"detalles"`
}

type DetallePedido struct {
	ID            int     `json:"id"`
	PedidoID      int     `json:"pedido_id"`
	ProductoID    int     `json:"producto_id"`
	Cantidad      int     `json:"cantidad"`
	PrecioUnitario float64 `json:"precio_unitario"`
	Subtotal      float64 `json:"subtotal"`
	Producto      *Producto `json:"producto,omitempty"`
}