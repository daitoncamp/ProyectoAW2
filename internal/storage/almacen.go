package storage

import "Proyecto_AWEBII/internal/models"

type EventoRepository interface {
	ListarEventos() []models.Evento
	BuscarEventoPorID(id int) (models.Evento, bool)
	CrearEvento(evento models.Evento) models.Evento
	ActualizarEvento(id int, evento models.Evento) (models.Evento, bool)
	BorrarEvento(id int) bool
}

type Almacen interface {
	EventoRepository
}

var _ Almacen = (*AlmacenSQLite)(nil)
var _ Almacen = (*Memoria)(nil)
