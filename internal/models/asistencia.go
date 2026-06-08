package models

import "time"

type Asistencia struct {
	ID           int       `json:"id"`
	EventoID     int       `json:"evento_id"`
	EstudianteID int       `json:"estudiante_id"`
	Presente     bool      `json:"presente"`
	HoraRegistro time.Time `json:"hora_registro"`
}
