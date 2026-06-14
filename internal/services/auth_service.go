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

	// Intentar con bcrypt primero
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		// Soporte para contraseña en texto plano del admin inicial
		if user.Password == password {
			// Hashear la contraseña para mejorar la seguridad
			hashed, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
			s.repo.UpdatePassword(user.ID, string(hashed)) // Actualizar a hash
		} else {
			return "", err
		}
	}

	// Generar token JWT
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"email":   user.Email,
		"rol_id":  user.RolID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	})

	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}