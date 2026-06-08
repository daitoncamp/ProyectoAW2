package models

type CategoriaEvento struct {
	ID          int    `json:"id"`
	Nombre      string `json:"nombre"`
	Descripcion string `json:"descripcion"`
}
