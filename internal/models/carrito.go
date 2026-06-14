package models

type Carrito struct {
	ID              int               `json:"id"`
	UsuarioID       int               `json:"usuario_id"`
	CarritoDetalles []CarritoDetalle  `json:"detalles"`
}

type CarritoDetalle struct {
	ID         int       `json:"id"`
	CarritoID  int       `json:"carrito_id"`
	ProductoID int       `json:"producto_id"`
	Cantidad   int       `json:"cantidad"`
	Producto   *Producto `json:"producto,omitempty"`
}