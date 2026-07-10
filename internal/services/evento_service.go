package services

import (
	"Proyecto_AWEBII/internal/models"
)

type EventoRepository interface {
	ListarEventos() []models.Evento
	BuscarEventoPorID(id int) (models.Evento, bool)
	CrearEvento(evento models.Evento) models.Evento
	ActualizarEvento(id int, evento models.Evento) (models.Evento, bool)
	BorrarEvento(id int) bool
}

type EventoService struct {
	repo EventoRepository
}

func NewEventoService(repo EventoRepository) *EventoService {
	return &EventoService{repo: repo}
}

func (s *EventoService) Listar() ([]models.Evento, error) {
	return s.repo.ListarEventos(), nil
}

func (s *EventoService) ObtenerPorID(id int) (models.Evento, error) {
	evento, ok := s.repo.BuscarEventoPorID(id)
	if !ok {
		return models.Evento{}, ErrEventoNoEncontrado
	}
	return evento, nil
}

func (s *EventoService) Crear(evento models.Evento) (models.Evento, error) {
	if evento.Nombre == "" {
		return models.Evento{}, ErrEventoNombreVacio
	}

	return s.repo.CrearEvento(evento), nil
}

func (s *EventoService) Actualizar(id int, evento models.Evento) (models.Evento, error) {
	if evento.Nombre == "" {
		return models.Evento{}, ErrEventoNombreVacio
	}

	updated, ok := s.repo.ActualizarEvento(id, evento)
	if !ok {
		return models.Evento{}, ErrEventoNoEncontrado
	}

	return updated, nil
}

func (s *EventoService) Eliminar(id int) error {
	ok := s.repo.BorrarEvento(id)
	if !ok {
		return ErrEventoNoEncontrado
	}
	return nil
}
