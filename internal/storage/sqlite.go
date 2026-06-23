package storage

import (
	"gorm.io/gorm"

	"Proyecto_AWEBII/internal/models"
)

// AlmacenSQLite implementa la interfaz Almacen usando GORM + SQLite.
type AlmacenSQLite struct {
	db *gorm.DB
}

// ActualizarEvento implements [Almacen].
func (a *AlmacenSQLite) ActualizarEvento(id int, datos models.Evento) (models.Evento, bool) {
	panic("unimplemented")
}

// BorrarEvento implements [Almacen].
func (a *AlmacenSQLite) BorrarEvento(id int) bool {
	panic("unimplemented")
}

// BuscarEventoPorID implements [Almacen].
func (a *AlmacenSQLite) BuscarEventoPorID(id int) (models.Evento, bool) {
	panic("unimplemented")
}

// CrearEvento implements [Almacen].
func (a *AlmacenSQLite) CrearEvento(e models.Evento) models.Evento {
	panic("unimplemented")
}

// ListarEventos implements [Almacen].
func (a *AlmacenSQLite) ListarEventos() []models.Evento {
	panic("unimplemented")
}

// NuevoAlmacenSQLite envuelve una conexión *gorm.DB ya abierta.
func NuevoAlmacenSQLite(db *gorm.DB) *AlmacenSQLite {
	return &AlmacenSQLite{
		db: db,
	}
}

// =========================================================
// INVERSIONES
// =========================================================

func (a *AlmacenSQLite) ListarInversiones() []models.Inversion {

	var inversiones []models.Inversion

	a.db.Find(&inversiones)

	return inversiones
}

func (a *AlmacenSQLite) BuscarInversionPorID(id int) (models.Inversion, bool) {

	var inversion models.Inversion

	if err := a.db.First(&inversion, id).Error; err != nil {
		return models.Inversion{}, false
	}

	return inversion, true
}

func (a *AlmacenSQLite) CrearInversion(i models.Inversion) models.Inversion {

	a.db.Create(&i)

	return i
}

func (a *AlmacenSQLite) ActualizarInversion(id int, datos models.Inversion) (models.Inversion, bool) {

	var existente models.Inversion

	if err := a.db.First(&existente, id).Error; err != nil {
		return models.Inversion{}, false
	}

	datos.ID = id

	a.db.Save(&datos)

	return datos, true
}

func (a *AlmacenSQLite) BorrarInversion(id int) bool {

	res := a.db.Delete(&models.Inversion{}, id)

	return res.RowsAffected > 0
}

// =========================================================
// SEEDS
// =========================================================

func (a *AlmacenSQLite) SembrarSiVacio() {

	var n int64

	a.db.Model(&models.Inversion{}).Count(&n)

	if n > 0 {
		return
	}

	inversiones := []models.Inversion{
		{
			ID:                  1,
			Nombre:              "Remodelación de Biblioteca",
			MontoInicial:        1000,
			MontoActual:         1100,
			RendimientoEsperado: 10,
			Estado:              "Activa",
			TipoInversionID:     1,
			DestinoInversionID:  1,
		},
		{
			ID:                  2,
			Nombre:              "Actualización de Equipos",
			MontoInicial:        500,
			MontoActual:         550,
			RendimientoEsperado: 10,
			Estado:              "Activa",
			TipoInversionID:     1,
			DestinoInversionID:  1,
		},
		{
			ID:                  3,
			Nombre:              "Proyecto de Investigación",
			MontoInicial:        2000,
			MontoActual:         2200,
			RendimientoEsperado: 10,
			Estado:              "Activa",
			TipoInversionID:     1,
			DestinoInversionID:  1,
		},
	}

	a.db.Create(&inversiones)
}

// Verificación en tiempo de compilación.
//var _ Almacen = (*AlmacenSQLite)(nil)
