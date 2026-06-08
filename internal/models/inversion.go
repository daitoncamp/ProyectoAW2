package models

// Inversion representa una inversión realizada por la asociación.
type Inversion struct {
	ID                  int     `json:"id"`
	Nombre              string  `json:"nombre"`
	MontoInicial        float64 `json:"monto_inicial"`
	MontoActual         float64 `json:"monto_actual"`
	RendimientoEsperado float64 `json:"rendimiento_esperado"`
	Estado              string  `json:"estado"`
	TipoInversionID     int     `json:"tipo_inversion_id"`
	DestinoInversionID  int     `json:"destino_inversion_id"`
}

// TipoInversion representa la categoría de una inversión.
type TipoInversion struct {
	ID          int    `json:"id"`
	Nombre      string `json:"nombre"`
	Descripcion string `json:"descripcion"`
}

// DestinoInversion representa el objetivo o destino de una inversión.
type DestinoInversion struct {
	ID          int    `json:"id"`
	Nombre      string `json:"nombre"`
	Descripcion string `json:"descripcion"`
}

// Aporte representa una contribución realizada a una inversión.
type Aporte struct {
	ID          int     `json:"id"`
	InversionID int     `json:"inversion_id"`
	Nombre      string  `json:"nombre"`
	Monto       float64 `json:"monto"`
}
