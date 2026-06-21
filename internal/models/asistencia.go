package models

import "time"

type Asistencia struct {
	ID           int       `gorm:"primaryKey;autoIncrement" json:"id"`
	EventoID     int       `json:"evento_id"`
	Presente     bool      `json:"presente"`
	HoraRegistro time.Time `json:"hora_registro"`
}
