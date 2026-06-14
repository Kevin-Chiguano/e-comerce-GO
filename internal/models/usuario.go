package models

import "time"

type Usuario struct {
	ID            int       `json:"id"`
	Nombre        string    `json:"nombre"`
	Email         string    `json:"email"`
	Password      string    `json:"-"`
	RolID         int       `json:"rol_id"`
	FechaRegistro time.Time `json:"fecha_registro"`
}