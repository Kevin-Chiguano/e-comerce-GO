package repositories

import (
	"database/sql"
	"ecommerce-manager/internal/models"
)

type ProductoRepository struct {
	db *sql.DB
}

func NewProductoRepository(db *sql.DB) *ProductoRepository {
	return &ProductoRepository{db: db}
}

func (r *ProductoRepository) GetAll() ([]models.Producto, error) {
	rows, err := r.db.Query(`
		SELECT id, nombre, descripcion, precio, stock, categoria_id, imagen 
		FROM productos`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var productos []models.Producto
	for rows.Next() {
		var p models.Producto
		err := rows.Scan(&p.ID, &p.Nombre, &p.Descripcion, &p.Precio, &p.Stock, &p.CategoriaID, &p.Imagen)
		if err != nil {
			return nil, err
		}
		productos = append(productos, p)
	}
	return productos, nil
}

func (r *ProductoRepository) GetByID(id int) (*models.Producto, error) {
	p := &models.Producto{}
	err := r.db.QueryRow(`
		SELECT id, nombre, descripcion, precio, stock, categoria_id, imagen 
		FROM productos WHERE id = $1`, id).
		Scan(&p.ID, &p.Nombre, &p.Descripcion, &p.Precio, &p.Stock, &p.CategoriaID, &p.Imagen)
	if err != nil {
		return nil, err
	}
	return p, nil
}