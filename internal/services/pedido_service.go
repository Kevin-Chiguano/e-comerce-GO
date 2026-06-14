package services

import (
	"database/sql"
	"ecommerce-manager/internal/models"
	"ecommerce-manager/internal/repositories"
)

type PedidoService struct {
	pedidoRepo  *repositories.PedidoRepository
	carritoRepo *repositories.CarritoRepository
}

func NewPedidoService(pedidoRepo *repositories.PedidoRepository, carritoRepo *repositories.CarritoRepository) *PedidoService {
	return &PedidoService{
		pedidoRepo:  pedidoRepo,
		carritoRepo: carritoRepo,
	}
}

func (s *PedidoService) CrearPedidoDesdeCarrito(usuarioID int) (*models.Pedido, error) {
	carrito, err := s.carritoRepo.GetOrCreate(usuarioID)
	if err != nil {
		return nil, err
	}

	detalles, err := s.carritoRepo.GetDetalles(carrito.ID)
	if err != nil {
		return nil, err
	}

	if len(detalles) == 0 {
		return nil, sql.ErrNoRows
	}

	// Calcular total y preparar detalles
	var total float64
	pedido := &models.Pedido{
		UsuarioID: usuarioID,
		Estado:    "PENDIENTE",
		Detalles:  make([]models.DetallePedido, 0),
	}

	for _, d := range detalles {
		subtotal := float64(d.Cantidad) * d.Producto.Precio
		total += subtotal

		detallePedido := models.DetallePedido{
			ProductoID:     d.ProductoID,
			Cantidad:       d.Cantidad,
			PrecioUnitario: d.Producto.Precio,
			Subtotal:       subtotal,
		}
		pedido.Detalles = append(pedido.Detalles, detallePedido)
	}

	pedido.Total = total

	// Crear el pedido
	err = s.pedidoRepo.Create(pedido)
	if err != nil {
		return nil, err
	}

	// Guardar los detalles del pedido
	for i := range pedido.Detalles {
		pedido.Detalles[i].PedidoID = pedido.ID
		err = s.pedidoRepo.AddDetalle(pedido.Detalles[i])
		if err != nil {
			return nil, err
		}
	}

	return pedido, nil
}