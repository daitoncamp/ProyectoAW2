package mocks

import "Proyecto_AWEBII/internal/models"

type MockEventoAlmacen struct {
	Eventos []models.Evento
}

func (m *MockEventoAlmacen) ListarEventos() []models.Evento {
	return m.Eventos
}

func (m *MockEventoAlmacen) BuscarEventoPorID(id int) (models.Evento, bool) {

	for _, e := range m.Eventos {
		if e.ID == id {
			return e, true
		}
	}

	return models.Evento{}, false
}

func (m *MockEventoAlmacen) CrearEvento(e models.Evento) models.Evento {

	e.ID = len(m.Eventos) + 1

	m.Eventos = append(m.Eventos, e)

	return e
}

func (m *MockEventoAlmacen) ActualizarEvento(id int, e models.Evento) (models.Evento, bool) {

	for i, evento := range m.Eventos {

		if evento.ID == id {

			e.ID = id
			m.Eventos[i] = e

			return e, true
		}
	}

	return models.Evento{}, false
}

func (m *MockEventoAlmacen) BorrarEvento(id int) bool {

	for i, e := range m.Eventos {

		if e.ID == id {

			m.Eventos = append(
				m.Eventos[:i],
				m.Eventos[i+1:]...,
			)

			return true
		}
	}

	return false
}
