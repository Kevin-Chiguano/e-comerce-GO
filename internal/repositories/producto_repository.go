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
		SELECT id, nombre, descripcion, precio, stock, categoria_id, 
		       COALESCE(imagen, '') as imagen 
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
		SELECT id, nombre, descripcion, precio, stock, categoria_id, 
		       COALESCE(imagen, '') as imagen 
		FROM productos WHERE id = $1`, id).
		Scan(&p.ID, &p.Nombre, &p.Descripcion, &p.Precio, &p.Stock, &p.CategoriaID, &p.Imagen)
	if err != nil {
		return nil, err
	}
	return p, nil
}

func (r *ProductoRepository) Create(producto *models.Producto) error {
	err := r.db.QueryRow(`
		INSERT INTO productos (nombre, descripcion, precio, stock, categoria_id, imagen)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id`,
		producto.Nombre, producto.Descripcion, producto.Precio, 
		producto.Stock, producto.CategoriaID, producto.Imagen).
		Scan(&producto.ID)
	return err
}

func (r *ProductoRepository) Update(producto *models.Producto) error {
	_, err := r.db.Exec(`
		UPDATE productos 
		SET nombre = $1, descripcion = $2, precio = $3, stock = $4, categoria_id = $5, imagen = $6
		WHERE id = $7`,
		producto.Nombre, producto.Descripcion, producto.Precio,
		producto.Stock, producto.CategoriaID, producto.Imagen, producto.ID)
	return err
}

func (r *ProductoRepository) Delete(id int) error {
	_, err := r.db.Exec("DELETE FROM productos WHERE id = $1", id)
	return err
}

func (r *ProductoRepository) UpdateStock(id, cantidad int) error {
	_, err := r.db.Exec(
		"UPDATE productos SET stock = stock + $1 WHERE id = $2",
		cantidad, id)
	return err
}