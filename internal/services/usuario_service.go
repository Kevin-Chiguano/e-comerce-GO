package services

import (
	"ecommerce-manager/internal/models"
	"ecommerce-manager/internal/repositories"
)

type UsuarioService struct {
	repo *repositories.UsuarioRepository
}

func NewUsuarioService(repo *repositories.UsuarioRepository) *UsuarioService {
	return &UsuarioService{repo: repo}
}

func (s *UsuarioService) GetByID(id int) (*models.Usuario, error) {
	// Puedes implementar si lo necesitas
	return nil, nil
}