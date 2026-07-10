package models

import "time"

type Evento struct {
	ID int `gorm:"primaryKey;autoIncrement" json:"id"`

	Nombre      string `gorm:"not null" json:"nombre"`
	Descripcion string `json:"descripcion"`
	Fecha       string `json:"fecha"`
	Lugar       string `json:"lugar"`
	Capacidad   int    `json:"capacidad"`
	CategoriaID int    `json:"categoria_id"`
	Organizador string `json:"organizador"`
	Estado      string `json:"estado"`
}

type CategoriaEvento struct {
	ID int `gorm:"primaryKey;autoIncrement" json:"id"`

	Nombre      string `gorm:"not null" json:"nombre"`
	Descripcion string `json:"descripcion"`
}
type Asistencia struct {
	ID           int       `gorm:"primaryKey;autoIncrement" json:"id"`
	EventoID     int       `json:"evento_id"`
	Presente     bool      `json:"presente"`
	HoraRegistro time.Time `json:"hora_registro"`
}
