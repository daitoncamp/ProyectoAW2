package storage

import (
	"gorm.io/gorm"

	"Proyecto_AWEBII/internal/models"
)

type AlmacenSQLite struct {
	db *gorm.DB
}

func NuevoAlmacenSQLite(db *gorm.DB) *AlmacenSQLite {
	return &AlmacenSQLite{
		db: db,
	}
}

// Eventos

func (a *AlmacenSQLite) ListarEventos() []models.Evento {

	var eventos []models.Evento
	a.db.
		// Preload("Categoria").
		// Preload("Asistencias").
		Find(&eventos)
	return eventos
}

func (a *AlmacenSQLite) BuscarEventoPorID(id int) (models.Evento, bool) {
	var evento models.Evento
	err := a.db.
		// Preload("Categoria").
		// Preload("Asistencias").
		First(&evento, id).Error
	if err != nil {
		return models.Evento{}, false
	}
	return evento, true
}

func (a *AlmacenSQLite) CrearEvento(e models.Evento) models.Evento {
	a.db.Create(&e)
	return e
}

func (a *AlmacenSQLite) ActualizarEvento(id int, datos models.Evento) (models.Evento, bool) {
	var existente models.Evento
	if err := a.db.First(&existente, id).Error; err != nil {
		return models.Evento{}, false
	}
	datos.ID = id
	a.db.Save(&datos)
	return datos, true
}

func (a *AlmacenSQLite) BorrarEvento(id int) bool {
	res := a.db.Delete(&models.Evento{}, id)

	return res.RowsAffected > 0
}

func (a *AlmacenSQLite) SembrarSiVacioEventos() {

	var n int64

	a.db.Model(&models.Evento{}).Count(&n)

	if n > 0 {
		return
	}

	eventos := []models.Evento{
		{
			Nombre:      "Taller de Go",
			Descripcion: "Introducción al lenguaje Go",
			Fecha:       "2026-06-20",
			Lugar:       "Laboratorio A",
			Capacidad:   50,
			CategoriaID: 1,
			Organizador: "Asociación Estudiantil",
			Estado:      "Activo",
		},
		{
			Nombre:      "Festival Cultural",
			Descripcion: "Presentaciones artísticas",
			Fecha:       "2026-06-25",
			Lugar:       "Auditorio Principal",
			Capacidad:   200,
			CategoriaID: 2,
			Organizador: "Departamento de Cultura",
			Estado:      "Activo",
		},
		{
			Nombre:      "Campeonato de Fútbol",
			Descripcion: "Torneo interfacultades",
			Fecha:       "2026-07-01",
			Lugar:       "Cancha Central",
			Capacidad:   150,
			CategoriaID: 3,
			Organizador: "Liga Deportiva Universitaria",
			Estado:      "Programado",
		},
	}

	a.db.Create(&eventos)
}


// Verificación de interfaz

var _ Almacen = (*AlmacenSQLite)(nil)
