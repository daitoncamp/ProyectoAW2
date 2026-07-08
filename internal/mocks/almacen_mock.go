package mocks

import (
	"Proyecto_AWEBII/internal/models"
)

type MockAlmacen struct {
	Inversiones []models.Inversion
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

afunc (m *MockAlmacen) BorrarInversion(id int) bool {

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

// ===============================
// Eventos (Mock vacío)
// ===============================

func (m *MockAlmacen) ListarEventos() []models.Evento {
	return nil
}

func (m *MockAlmacen) BuscarEventoPorID(id int) (models.Evento, bool) {
	return models.Evento{}, false
}

func (m *MockAlmacen) CrearEvento(e models.Evento) models.Evento {
	return models.Evento{}
}

func (m *MockAlmacen) ActualizarEvento(id int, datos models.Evento) (models.Evento, bool) {
	return models.Evento{}, false
}

func (m *MockAlmacen) BorrarEvento(id int) bool {
	return false
}
