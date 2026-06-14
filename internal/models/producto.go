package models

type Producto struct {
	ID          int     `json:"id"`
	Nombre      string  `json:"nombre"`
	Descripcion string  `json:"descripcion"`
	Precio      float64 `json:"precio"`
	Stock       int     `json:"stock"`
	CategoriaID int     `json:"categoria_id"`
	Imagen      string  `json:"imagen"`
}