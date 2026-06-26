package repositories

import (
	"database/sql"
	"ecommerce-manager/internal/models"
)

type PedidoRepository struct {
	db *sql.DB
}

func NewPedidoRepository(db *sql.DB) *PedidoRepository {
	return &PedidoRepository{db: db}
}

func (r *PedidoRepository) Create(pedido *models.Pedido) error {
	err := r.db.QueryRow(`
		INSERT INTO pedidos (usuario_id, total, estado) 
		VALUES ($1, $2, $3) RETURNING id, fecha`,
		pedido.UsuarioID, pedido.Total, pedido.Estado).
		Scan(&pedido.ID, &pedido.Fecha)
	return err
}

func (r *PedidoRepository) AddDetalle(detalle models.DetallePedido) error {
	_, err := r.db.Exec(`
		INSERT INTO detalle_pedidos (pedido_id, producto_id, cantidad, precio_unitario, subtotal)
		VALUES ($1, $2, $3, $4, $5)`,
		detalle.PedidoID, detalle.ProductoID, detalle.Cantidad,
		detalle.PrecioUnitario, detalle.Subtotal)
	return err
}

func (r *PedidoRepository) GetByUsuario(usuarioID int) ([]models.Pedido, error) {
	rows, err := r.db.Query(`
		SELECT id, usuario_id, total, estado, fecha 
		FROM pedidos WHERE usuario_id = $1 ORDER BY fecha DESC`, usuarioID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var pedidos []models.Pedido
	for rows.Next() {
		var p models.Pedido
		err := rows.Scan(&p.ID, &p.UsuarioID, &p.Total, &p.Estado, &p.Fecha)
		if err != nil {
			return nil, err
		}
		pedidos = append(pedidos, p)
	}
	return pedidos, nil
}

func (r *PedidoRepository) GetAll() ([]models.Pedido, error) {
	rows, err := r.db.Query(`
		SELECT id, usuario_id, total, estado, fecha
		FROM pedidos ORDER BY fecha DESC`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var pedidos []models.Pedido
	for rows.Next() {
		var p models.Pedido
		err := rows.Scan(&p.ID, &p.UsuarioID, &p.Total, &p.Estado, &p.Fecha)
		if err != nil {
			return nil, err
		}
		pedidos = append(pedidos, p)
	}
	return pedidos, nil
}

func (r *PedidoRepository) UpdateEstado(pedidoID int, estado string) error {
	_, err := r.db.Exec(`UPDATE pedidos SET estado = $1 WHERE id = $2`, estado, pedidoID)
	return err
}