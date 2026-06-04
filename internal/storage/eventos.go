package eventos

import "time"

type Evento struct {
	ID uint `gorm:"primaryKey" json:"id"`

	Nombre      string `gorm:"size:100;not null" json:"nombre"`
	Descripcion string `gorm:"type:text" json:"descripcion"`

	Tipo  string `gorm:"size:50" json:"tipo"`
	Lugar string `gorm:"size:100" json:"lugar"`

	Fecha time.Time `json:"fecha"`

	Capacidad int `gorm:"not null" json:"capacidad"`

	CuposDisponibles int `gorm:"not null" json:"cupos_disponibles"`

	Estado string `gorm:"size:30;default:'activo'" json:"estado"`

	InversionID uint `json:"inversion_id"`

	CreadoEn      time.Time `json:"creado_en"`
	ActualizadoEn time.Time `json:"actualizado_en"`

	Participantes []ParticipacionEvento `gorm:"foreignKey:EventoID" json:"participantes"`
}

type ParticipacionEvento struct {
	ID uint `gorm:"primaryKey" json:"id"`

	EventoID     uint `json:"evento_id"`
	EstudianteID uint `json:"estudiante_id"`

	FechaRegistro time.Time `json:"fecha_registro"`

	Asistencia bool `gorm:"default:false" json:"asistencia"`

	Estado string `gorm:"size:20;default:'registrado'" json:"estado"`
}
