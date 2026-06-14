package services

import (
	"ecommerce-manager/internal/models"
	"ecommerce-manager/internal/repositories"
)

type CarritoService struct {
	repo *repositories.CarritoRepository
}

func NewCarritoService(repo *repositories.CarritoRepository) *CarritoService {
	return &CarritoService{repo: repo}
}

func (s *CarritoService) GetOrCreateCarrito(usuarioID int) (*models.Carrito, error) {
	carrito, err := s.repo.GetOrCreate(usuarioID)
	if err != nil {
		return nil, err
	}

	detalles, err := s.repo.GetDetalles(carrito.ID)
	if err != nil {
		return nil, err
	}
	carrito.CarritoDetalles = detalles

	return carrito, nil
}

func (s *CarritoService) AddToCarrito(usuarioID, productoID, cantidad int) error {
	carrito, err := s.repo.GetOrCreate(usuarioID)
	if err != nil {
		return err
	}
	return s.repo.AddItem(carrito.ID, productoID, cantidad)
}

func (s *CarritoService) RemoveFromCarrito(usuarioID, productoID int) error {
	carrito, err := s.repo.GetOrCreate(usuarioID)
	if err != nil {
		return err
	}
	return s.repo.RemoveItem(carrito.ID, productoID)
}

func (s *CarritoService) ClearCarrito(usuarioID int) error {
	carrito, err := s.repo.GetOrCreate(usuarioID)
	if err != nil {
		return err
	}
	return s.repo.Clear(carrito.ID)
}