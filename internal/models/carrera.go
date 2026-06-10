package models

type Carrera struct {
	ID          int    `json:"id"`
	Nombre      string `json:"nombre"`
	Facultad    string `json:"facultad"`
	Descripcion string `json:"descripcion"`
}
