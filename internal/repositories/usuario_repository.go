package repositories

import (
	"database/sql"
	"ecommerce-manager/internal/models"
)

type UsuarioRepository struct {
	db *sql.DB
}

func NewUsuarioRepository(db *sql.DB) *UsuarioRepository {
	return &UsuarioRepository{db: db}
}

func (r *UsuarioRepository) Create(user *models.Usuario) error {
	return r.db.QueryRow("INSERT INTO usuarios (nombre, email, password, rol_id) VALUES ($1, $2, $3, $4) RETURNING id",
		user.Nombre, user.Email, user.Password, user.RolID).Scan(&user.ID)
}

func (r *UsuarioRepository) GetByEmail(email string) (*models.Usuario, error) {
	user := &models.Usuario{}
	err := r.db.QueryRow("SELECT id, nombre, email, password, rol_id, fecha_registro FROM usuarios WHERE email = $1", email).
		Scan(&user.ID, &user.Nombre, &user.Email, &user.Password, &user.RolID, &user.FechaRegistro)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// Add more methods as needed...