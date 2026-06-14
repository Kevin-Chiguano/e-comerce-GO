package services

import (
	"ecommerce-manager/internal/models"
	"ecommerce-manager/internal/repositories"
	"golang.org/x/crypto/bcrypt"
	"github.com/golang-jwt/jwt/v5"
	"os"
	"time"
	"fmt"
)

type AuthService struct {
	repo *repositories.UsuarioRepository
}

func NewAuthService(repo *repositories.UsuarioRepository) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) Register(user *models.Usuario) error {

	fmt.Println("Nombre:", user.Nombre)
	fmt.Println("Email:", user.Email)
	fmt.Println("Password recibida:", user.Password)

	hashed, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user.Password = string(hashed)

	fmt.Println("Password encriptada:", user.Password)

	return s.repo.Create(user)
}

func (s *AuthService) Login(email, password string) (string, error) {

	fmt.Println("Email recibido:", email)
	fmt.Println("Password recibida:", password)

	user, err := s.repo.GetByEmail(email)
	if err != nil {
		fmt.Println("Usuario no encontrado")
		return "", err
	}

	fmt.Println("Password guardada en BD:", user.Password)

	err = bcrypt.CompareHashAndPassword(
		[]byte(user.Password),
		[]byte(password),
	)

	if err == nil {
		fmt.Println("Login correcto")
	} else if user.Password == password {
		hashed, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		s.repo.UpdatePassword(user.ID, string(hashed))
	} else {
		fmt.Println("Contraseña incorrecta")
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"email":   user.Email,
		"rol_id":  user.RolID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	})

	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}