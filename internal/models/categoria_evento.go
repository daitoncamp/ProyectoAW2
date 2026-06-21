package models

type CategoriaEvento struct {
	ID int `gorm:"primaryKey;autoIncrement" json:"id"`

	Nombre      string `gorm:"not null" json:"nombre"`
	Descripcion string `json:"descripcion"`
}
