package services

import (
	"Proyecto_AWEBII/internal/models"
	"Proyecto_AWEBII/internal/storage"
)

type EventoService struct {
	repo storage.Almacen
}

func NewEventoService(repo storage.Almacen) *EventoService {
	return &EventoService{repo: repo}
}

func (s *EventoService) ObtenerEventos() []models.Evento {
	return s.repo.ListarEventos()
}

func (s *EventoService) ObtenerEventoPorID(id int) (models.Evento, bool) {
	return s.repo.BuscarEventoPorID(id)
}

func (s *EventoService) CrearEvento(evento models.Evento) models.Evento {
	return s.repo.CrearEvento(evento)
}

func (s *EventoService) ActualizarEvento(id int, evento models.Evento) (models.Evento, bool) {
	return s.repo.ActualizarEvento(id, evento)
}

func (s *EventoService) EliminarEvento(id int) bool {
	return s.repo.BorrarEvento(id)
}
