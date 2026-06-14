package repositories

import (
	"database/sql"
	"ecommerce-manager/internal/models"
)

type CarritoRepository struct {
	db *sql.DB
}

func NewCarritoRepository(db *sql.DB) *CarritoRepository {
	return &CarritoRepository{db: db}
}

// Obtener o crear carrito del usuario
func (r *CarritoRepository) GetOrCreate(usuarioID int) (*models.Carrito, error) {
	carrito := &models.Carrito{UsuarioID: usuarioID}

	// Buscar si existe
	err := r.db.QueryRow("SELECT id FROM carrito WHERE usuario_id = $1", usuarioID).
		Scan(&carrito.ID)

	if err == sql.ErrNoRows {
		// Crear nuevo carrito
		err = r.db.QueryRow("INSERT INTO carrito (usuario_id) VALUES ($1) RETURNING id", usuarioID).
			Scan(&carrito.ID)
		if err != nil {
			return nil, err
		}
	} else if err != nil {
		return nil, err
	}

	return carrito, nil
}

func (r *CarritoRepository) AddItem(carritoID, productoID, cantidad int) error {
	_, err := r.db.Exec(`
		INSERT INTO carrito_detalle (carrito_id, producto_id, cantidad)
		VALUES ($1, $2, $3)
		ON CONFLICT (carrito_id, producto_id) 
		DO UPDATE SET cantidad = carrito_detalle.cantidad + $3`,
		carritoID, productoID, cantidad)
	return err
}

func (r *CarritoRepository) GetDetalles(carritoID int) ([]models.CarritoDetalle, error) {
	rows, err := r.db.Query(`
		SELECT cd.id, cd.carrito_id, cd.producto_id, cd.cantidad, 
		       p.id, p.nombre, p.precio, p.imagen
		FROM carrito_detalle cd
		JOIN productos p ON cd.producto_id = p.id
		WHERE cd.carrito_id = $1`, carritoID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var detalles []models.CarritoDetalle
	for rows.Next() {
		var d models.CarritoDetalle
		var p models.Producto
		err := rows.Scan(&d.ID, &d.CarritoID, &d.ProductoID, &d.Cantidad,
			&p.ID, &p.Nombre, &p.Precio, &p.Imagen)
		if err != nil {
			return nil, err
		}
		d.Producto = &p
		detalles = append(detalles, d)
	}
	return detalles, nil
}