package storage

import "Proyecto_AWEBII/internal/models"

type Almacen interface {
	//Inversiones
	ListarInversiones() []models.Inversion
	BuscarInversionPorID(id int) (models.Inversion, bool)
	CrearInversion(i models.Inversion) models.Inversion
	ActualizarInversion(id int, datos models.Inversion) (models.Inversion, bool)
	BorrarInversion(id int) bool

	// EVENTOS

	ListarEventos() []models.Evento
	BuscarEventoPorID(id int) (models.Evento, bool)
	CrearEvento(e models.Evento) models.Evento
	ActualizarEvento(id int, datos models.Evento) (models.Evento, bool)
	BorrarEvento(id int) bool
}

type UserRepository interface {
	CrearUsuario(u models.Usuario) (models.Usuario, error)
	BuscarUsuarioPorEmail(email string) (models.Usuario, bool)
}

//var _ Almacen = (*Memoria)(nil)
