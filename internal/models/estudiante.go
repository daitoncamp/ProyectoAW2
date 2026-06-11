package models

import "time"

type Estudiante struct {
	ID              int       `json:"id"`
	Nombres         string    `json:"nombres"`
	Apellidos       string    `json:"apellidos"`
	Email           string    `json:"email"`
	Telefono        string    `json:"telefono"`
	FechaNacimiento time.Time `json:"fecha_nacimiento"`
	CarreraID       int       `json:"carrera_id"`
	FechaRegistro   time.Time `json:"fecha_registro"`
}
