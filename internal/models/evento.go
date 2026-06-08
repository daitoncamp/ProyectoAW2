package models

import "time"

type Evento struct {
	ID          int       `json:"id"`
	Nombre      string    `json:"nombre"`
	Descripcion string    `json:"descripcion"`
	Fecha       time.Time `json:"fecha"`
	Lugar       string    `json:"lugar"`
	Capacidad   int       `json:"capacidad"`
	CategoriaID int       `json:"categoria_id"`
	Organizador string    `json:"organizador"`
	Estado      string    `json:"estado"`
}
