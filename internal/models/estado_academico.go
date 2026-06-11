package models

import "time"

type EstadoAcademico struct {
	ID                 int       `json:"id"`
	EstudianteID       int       `json:"estudiante_id"`
	Promedio           float64   `json:"promedio"`
	Asistencia         float64   `json:"asistencia"`
	Estado             string    `json:"estado"`
	FechaActualizacion time.Time `json:"fecha_actualizacion"`
}
