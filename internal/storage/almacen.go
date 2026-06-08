package storage

import "Proyecto_AWEBII/internal/models"

type Almacen interface {
	//Inversiones
	ListarInversiones() []models.Inversion
	BuscarInversionPorID(id int) (models.Inversion, bool)
	CrearInversion(i models.Inversion) models.Inversion
	ActualizarInversion(id int, datos models.Inversion) (models.Inversion, bool)
	BorrarInversion(id int) bool
}
