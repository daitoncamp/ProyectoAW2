package inversiones

import "time"

type Inversion struct {
	ID                  uint    `gorm:"primaryKey" json:"id"`
	Nombre              string  `gorm:"size:100;not null" json:"nombre"`
	Descripcion         string  `gorm:"type:text" json:"descripcion"`
	Tipo                string  `gorm:"size:50;not null" json:"tipo"`
	MontoInicial        float64 `gorm:"not null" json:"monto_inicial"`
	MontoActual         float64 `gorm:"not null" json:"monto_actual"`
	RendimientoEsperado float64 `gorm:"not null" json:"rendimiento_esperado"`
	Estado              string  `gorm:"size:30;default:'activa'" json:"estado"`

	FechaInicio time.Time `json:"fecha_inicio"`
	FechaFin    time.Time `json:"fecha_fin"`

	CreadoEn      time.Time `json:"creado_en"`
	ActualizadoEn time.Time `json:"actualizado_en"`

	Aportes []Aporte `gorm:"foreignKey:InversionID" json:"aportes"`
	Eventos []Evento `gorm:"foreignKey:InversionID" json:"eventos"`
}

type Aporte struct {
	ID uint `gorm:"primaryKey" json:"id"`

	EstudianteID uint `json:"estudiante_id"`
	InversionID  uint `json:"inversion_id"`

	Monto       float64   `gorm:"not null" json:"monto"`
	FechaAporte time.Time `json:"fecha_aporte"`

	EstadoPago string `gorm:"size:20;default:'pagado'" json:"estado_pago"`

	CreadoEn time.Time `json:"creado_en"`
}

type Evento struct {
	ID          uint `gorm:"primaryKey"`
	Nombre      string
	Descripcion string

	InversionID uint

	Fecha     time.Time
	Capacidad int
	Estado    string
}
