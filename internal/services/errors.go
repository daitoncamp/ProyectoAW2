package services

import "errors"

var (
	ErrCamposVacios          = errors.New("email y contraseña son obligatorios")
	ErrEmailEnUso            = errors.New("el correo ya está registrado")
	ErrCredencialesInvalidas = errors.New("credenciales inválidas")
)
