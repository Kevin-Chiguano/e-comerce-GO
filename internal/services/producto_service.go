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