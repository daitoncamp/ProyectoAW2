package storage

import (
	"sync"

	"Proyecto_AWEBII/internal/models"
)

type Memoria struct {
	inversiones []models.Inversion
	nextID      int

	eventos      []models.Evento
	nextEventoID int

	mu sync.Mutex
}

func NuevaMemoria() *Memoria {
	return &Memoria{
		inversiones:  []models.Inversion{},
		nextID:       1,
		eventos:      []models.Evento{},
		nextEventoID: 1,
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

// =========================================================
// EVENTOS
// =========================================================

func (m *Memoria) SeedEventos() {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.eventos = []models.Evento{
		{ID: 1, Nombre: "Taller de Go", Descripcion: "Introducción al lenguaje Go", Fecha: "2026-06-20", Lugar: "Laboratorio A", Capacidad: 50, CategoriaID: 1, Organizador: "Asociación Estudiantil", Estado: "Activo"},
		{ID: 2, Nombre: "Festival Cultural", Descripcion: "Presentaciones artísticas", Fecha: "2026-06-25", Lugar: "Auditorio Principal", Capacidad: 200, CategoriaID: 2, Organizador: "Departamento de Cultura", Estado: "Activo"},
		{ID: 3, Nombre: "Campeonato de Fútbol", Descripcion: "Torneo interfacultades", Fecha: "2026-07-01", Lugar: "Cancha Central", Capacidad: 150, CategoriaID: 3, Organizador: "Liga Deportiva Universitaria", Estado: "Programado"},
	}

	m.nextEventoID = 4
}

func (m *Memoria) ListarEventos() []models.Evento {
	m.mu.Lock()
	defer m.mu.Unlock()

	copia := make([]models.Evento, len(m.eventos))
	copy(copia, m.eventos)

	return copia
}

func (m *Memoria) BuscarEventoPorID(id int) (models.Evento, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()

	for _, e := range m.eventos {
		if e.ID == id {
			return e, true
		}
	}

	return models.Evento{}, false
}

func (m *Memoria) CrearEvento(e models.Evento) models.Evento {
	m.mu.Lock()
	defer m.mu.Unlock()

	e.ID = m.nextEventoID
	m.nextEventoID++

	m.eventos = append(m.eventos, e)

	return e
}

func (m *Memoria) ActualizarEvento(id int, datos models.Evento) (models.Evento, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()

	for idx, evento := range m.eventos {

		if evento.ID == id {

			datos.ID = id
			m.eventos[idx] = datos

			return datos, true
		}
	}

	return models.Evento{}, false
}

func (m *Memoria) BorrarEvento(id int) bool {
	m.mu.Lock()
	defer m.mu.Unlock()

	for idx, evento := range m.eventos {

		if evento.ID == id {

			m.eventos = append(
				m.eventos[:idx],
				m.eventos[idx+1:]...,
			)

			return true
		}
	}

	return false
}
