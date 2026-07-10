package mocks

import (
	"Proyecto_AWEBII/internal/models"
)

type MockAlmacen struct {
	Inversiones []models.Inversion
	Eventos     []models.Evento
}

// ===============================
// Inversiones
// ===============================

func (m *MockAlmacen) ListarInversiones() []models.Inversion {
	return m.Inversiones
}

func (m *MockAlmacen) BuscarInversionPorID(id int) (models.Inversion, bool) {

	for _, i := range m.Inversiones {
		if i.ID == id {
			return i, true
		}
	}

	return models.Inversion{}, false
}

func (m *MockAlmacen) CrearInversion(i models.Inversion) models.Inversion {

	i.ID = len(m.Inversiones) + 1

	m.Inversiones = append(m.Inversiones, i)

	return i
}

func (m *MockAlmacen) ActualizarInversion(id int, datos models.Inversion) (models.Inversion, bool) {

	for indice, inv := range m.Inversiones {

		if inv.ID == id {

			datos.ID = id

			m.Inversiones[indice] = datos

			return datos, true
		}
	}

	return models.Inversion{}, false
}

func (m *MockAlmacen) BorrarInversion(id int) bool {

	for indice, inv := range m.Inversiones {

		if inv.ID == id {

			m.Inversiones = append(
				m.Inversiones[:indice],
				m.Inversiones[indice+1:]...,
			)

			return true
		}
	}

	return false
}

// Eventos

func (m *MockAlmacen) ListarEventos() []models.Evento {
	return m.Eventos
}

func (m *MockAlmacen) BuscarEventoPorID(id int) (models.Evento, bool) {

	for _, e := range m.Eventos {
		if e.ID == id {
			return e, true
		}
	}

	return models.Evento{}, false
}

func (m *MockAlmacen) CrearEvento(e models.Evento) models.Evento {

	e.ID = len(m.Eventos) + 1

	m.Eventos = append(m.Eventos, e)

	return e
}

func (m *MockAlmacen) ActualizarEvento(id int, e models.Evento) (models.Evento, bool) {

	for i, evento := range m.Eventos {

		if evento.ID == id {

			e.ID = id
			m.Eventos[i] = e

			return e, true
		}
	}

	return models.Evento{}, false
}

func (m *MockAlmacen) BorrarEvento(id int) bool {

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
