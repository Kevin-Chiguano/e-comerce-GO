package services

import (
	"ecommerce-manager/internal/models"
	"ecommerce-manager/internal/repositories"
)

type ProductoService struct {
	repo *repositories.ProductoRepository
}

func NewProductoService(repo *repositories.ProductoRepository) *ProductoService {
	return &ProductoService{repo: repo}
}

func (s *ProductoService) GetAllProductos() ([]models.Producto, error) {
	return s.repo.GetAll()
}

func (s *ProductoService) GetProductoByID(id int) (*models.Producto, error) {
	return s.repo.GetByID(id)
}

func (s *ProductoService) CreateProducto(producto *models.Producto) error {
	return s.repo.Create(producto)
}

func (s *ProductoService) UpdateProducto(producto *models.Producto) error {
	return s.repo.Update(producto)
}

func (s *ProductoService) DeleteProducto(id int) error {
	return s.repo.Delete(id)
}

func (s *ProductoService) UpdateStock(productoID, cantidad int) error {
	return s.repo.UpdateStock(productoID, cantidad)
}