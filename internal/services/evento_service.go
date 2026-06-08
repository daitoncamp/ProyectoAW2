package services

import (
	"Proyecto_AWEBII/internal/models"
	"Proyecto_AWEBII/internal/storage"
	"errors"
)

func ObtenerEventos() []models.Evento {
	return storage.Eventos
}
func CrearEvento(evento models.Evento) {
	storage.Eventos = append(storage.Eventos, evento)
}

func ObtenerEventoPorID(id int) (models.Evento, error) {
	for _, evento := range storage.Eventos {
		if evento.ID == id {
			return evento, nil
		}
	}
	return models.Evento{}, errors.New("evento no encontrado")
}

func ActualizarEvento(id int, eventoActualizado models.Evento) error {
	for i, evento := range storage.Eventos {
		if evento.ID == id {
			storage.Eventos[i] = eventoActualizado
			return nil
		}
	}
	return errors.New("evento no encontrado")
}

func EliminarEvento(id int) error {
	for i, evento := range storage.Eventos {
		if evento.ID == id {
			storage.Eventos = append(storage.Eventos[:i], storage.Eventos[i+1:]...)
			return nil
		}
	}
	return errors.New("evento no encontrado")
}

//funciones CRUD

// CrearEvento()
// ObtenerEventos()
// ObtenerEventoPorID()
// ActualizarEvento()
// EliminarEvento()
