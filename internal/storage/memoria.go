package storage

import (
	"sync"

	"Proyecto_AWEBII/internal/models"
)

type Memoria struct {
	inversiones []models.Inversion
	nextID      int
	mu          sync.Mutex
}

func NuevaMemoria() *Memoria {
	return &Memoria{
		inversiones: []models.Inversion{},
		nextID:      1,
	}
}

// =========================================================
// INVERSIONES
// =========================================================
func (m *Memoria) SeedInversiones() {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.inversiones = []models.Inversion{
		{ID: 1, Nombre: "Juan Pérez", MontoInicial: 1000, MontoActual: 1100, RendimientoEsperado: 10, Estado: "activa", TipoInversionID: 1, DestinoInversionID: 1},
		{ID: 2, Nombre: "María Gómez", MontoInicial: 500, MontoActual: 550, RendimientoEsperado: 10, Estado: "activa", TipoInversionID: 1, DestinoInversionID: 1},
		{ID: 3, Nombre: "Carlos Rodríguez", MontoInicial: 2000, MontoActual: 2200, RendimientoEsperado: 10, Estado: "activa", TipoInversionID: 1, DestinoInversionID: 1},
		{ID: 4, Nombre: "Ana Martínez", MontoInicial: 1500, MontoActual: 1650, RendimientoEsperado: 10, Estado: "activa", TipoInversionID: 1, DestinoInversionID: 1},
		{ID: 5, Nombre: "Luis Fernández", MontoInicial: 800, MontoActual: 880, RendimientoEsperado: 10, Estado: "activa", TipoInversionID: 1, DestinoInversionID: 1},
		{ID: 6, Nombre: "Sofía López", MontoInicial: 1200, MontoActual: 1320, RendimientoEsperado: 10, Estado: "activa", TipoInversionID: 1, DestinoInversionID: 1},
		{ID: 7, Nombre: "Miguel Sánchez", MontoInicial: 700, MontoActual: 770, RendimientoEsperado: 10, Estado: "activa", TipoInversionID: 1, DestinoInversionID: 1},
	}
	m.nextID = 8
}

// ListarInversiones devuelve todas las inversiones en memoria.

func (m *Memoria) ListarInversiones() []models.Inversion {
	m.mu.Lock()
	defer m.mu.Unlock()

	copia := make([]models.Inversion, len(m.inversiones))
	copy(copia, m.inversiones)

	return copia
}

func (m *Memoria) BuscarInversionPorID(id int) (models.Inversion, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()

	for _, i := range m.inversiones {
		if i.ID == id {
			return i, true
		}
	}

	return models.Inversion{}, false
}

func (m *Memoria) CrearInversion(i models.Inversion) models.Inversion {
	m.mu.Lock()
	defer m.mu.Unlock()

	i.ID = m.nextID
	m.nextID++

	m.inversiones = append(m.inversiones, i)

	return i
}

func (m *Memoria) ActualizarInversion(id int, datos models.Inversion) (models.Inversion, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()

	for idx, inv := range m.inversiones {

		if inv.ID == id {

			datos.ID = id
			m.inversiones[idx] = datos

			return datos, true
		}
	}

	return models.Inversion{}, false
}

func (m *Memoria) BorrarInversion(id int) bool {
	m.mu.Lock()
	defer m.mu.Unlock()

	for idx, inv := range m.inversiones {

		if inv.ID == id {

			m.inversiones = append(
				m.inversiones[:idx],
				m.inversiones[idx+1:]...,
			)

			return true
		}
	}

	return false
}
