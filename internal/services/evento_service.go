package services

import (
	"errors"

	"Proyecto_AWEBII/internal/models"
	"Proyecto_AWEBII/internal/storage"
)

var (
	ErrNoEncontrado = errors.New("evento no encontrado")
	ErrNombreVacio  = errors.New("nombre vacío")
)

type EventoService struct {
	repo storage.Almacen
}

func NewEventoService(repo storage.Almacen) *EventoService {
	return &EventoService{repo: repo}
}

func (s *EventoService) Listar() ([]models.Evento, error) {
	return s.repo.ListarEventos(), nil
}
func (s *EventoService) ObtenerPorID(id int) (models.Evento, error) {
	evento, ok := s.repo.BuscarEventoPorID(id)
	if !ok {
		return models.Evento{}, ErrNoEncontrado
	}
	return evento, nil
}
func (s *EventoService) Crear(evento models.Evento) (models.Evento, error) {
	if evento.Nombre == "" {
		return models.Evento{}, ErrNombreVacio
	}

	return s.repo.CrearEvento(evento), nil
}
func (s *EventoService) Actualizar(id int, evento models.Evento) (models.Evento, error) {
	if evento.Nombre == "" {
		return models.Evento{}, ErrNombreVacio
	}

	updated, ok := s.repo.ActualizarEvento(id, evento)
	if !ok {
		return models.Evento{}, ErrNoEncontrado
	}

	return updated, nil
}
func (s *EventoService) Eliminar(id int) error {
	ok := s.repo.BorrarEvento(id)
	if !ok {
		return ErrNoEncontrado
	}
	return nil
}
