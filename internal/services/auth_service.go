package services

import (
	"ecommerce-manager/internal/models"
	"ecommerce-manager/internal/repositories"
	"golang.org/x/crypto/bcrypt"
	"github.com/golang-jwt/jwt/v5"
	"os"
	"time"
)

type AuthService struct {
	repo *repositories.UsuarioRepository
}

func NewAuthService(repo *repositories.UsuarioRepository) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) Register(user *models.Usuario) error {
	hashed, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashed)
	return s.repo.Create(user)
}

func (s *AuthService) Login(email, password string) (string, error) {
	user, err := s.repo.GetByEmail(email)
	if err != nil {
		return "", err
	}

	// Intentar con bcrypt
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err == nil {
		// Contraseña correcta
	} else if user.Password == password {
		// Soporte para admin (contraseña en texto plano)
		// Actualizamos a hash para la próxima vez
		hashed, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		s.repo.UpdatePassword(user.ID, string(hashed))
	} else {
		return "", err
	}

	// Generar token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"email":   user.Email,
		"rol_id":  user.RolID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	})

	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}