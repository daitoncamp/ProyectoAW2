package storage

import (
	"sync"

	"Proyecto_AWEBII/internal/models"
)

type Memoria struct {
	eventos      []models.Evento
	nextEventoID int
	mu           sync.Mutex
}

func NuevaMemoria() *Memoria {
	return &Memoria{
		eventos:      []models.Evento{},
		nextEventoID: 1,
	}
}

// EVENTOSs
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

	for i, e := range m.eventos {
		if e.ID == id {
			datos.ID = id
			m.eventos[i] = datos
			return datos, true
		}
	}
	return models.Evento{}, false
}

func (m *Memoria) BorrarEvento(id int) bool {
	m.mu.Lock()
	defer m.mu.Unlock()

	for i, e := range m.eventos {
		if e.ID == id {
			m.eventos = append(m.eventos[:i], m.eventos[i+1:]...)
			return true
		}
	}
	return false
}
