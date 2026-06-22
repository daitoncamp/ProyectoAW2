package models

// Inversion representa una inversión realizada por la asociación.
type Inversion struct {
	ID                  int     `gorm:"primaryKey" json:"id"`
	Nombre              string  `gorm:"not null" json:"nombre"`
	MontoInicial        float64 `gorm:"not null" json:"monto_inicial"`
	MontoActual         float64 `json:"monto_actual"`
	RendimientoEsperado float64 `json:"rendimiento_esperado"`
	Estado              string  `gorm:"size:50;not null" json:"estado"`
	TipoInversionID     int     `json:"tipo_inversion_id"`
	DestinoInversionID  int     `json:"destino_inversion_id"`
}

// TipoInversion representa la categoría de una inversión.
type TipoInversion struct {
	ID          int    `gorm:"primaryKey" json:"id"`
	Nombre      string `gorm:"not null" json:"nombre"`
	Descripcion string `json:"descripcion"`
}

// DestinoInversion representa el objetivo o destino de una inversión.
type DestinoInversion struct {
	ID          int    `gorm:"primaryKey" json:"id"`
	Nombre      string `gorm:"not null" json:"nombre"`
	Descripcion string `json:"descripcion"`
}

// Aporte representa una contribución realizada a una inversión.
type Aporte struct {
	ID          int     `gorm:"primaryKey" json:"id"`
	InversionID int     `json:"inversion_id"`
	Nombre      string  `gorm:"not null" json:"nombre"`
	Monto       float64 `gorm:"not null" json:"mont"`
}
