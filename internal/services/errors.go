package services

import "errors"

// ==========================================
// Autenticación
// ==========================================

var (
	ErrCamposVacios          = errors.New("email y contraseña son obligatorios")
	ErrEmailEnUso            = errors.New("el correo ya está registrado")
	ErrCredencialesInvalidas = errors.New("credenciales inválidas")

	// ==========================================
	// Inversiones
	// ==========================================

	ErrNombreVacio   = errors.New("el nombre es obligatorio")
	ErrMontoInvalido = errors.New("el monto inicial debe ser mayor a 0")
	ErrEstadoVacio   = errors.New("el estado es obligatorio")
	ErrNoEncontrado  = errors.New("inversión no encontrada")
)

// Eventos

var (
	ErrEventoNombreVacio  = errors.New("el nombre del evento es obligatorio")
	ErrEventoNoEncontrado = errors.New("evento no encontrado")
)
