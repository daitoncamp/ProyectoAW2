package services

import (
	"Proyecto_AWEBII/internal/models"
	"Proyecto_AWEBII/internal/storage"
)

type InversionService struct {
	store storage.Almacen
}

func NewInversionService(store storage.Almacen) *InversionService {
	return &InversionService{
		store: store,
	}
}

func (s *InversionService) ListarInversiones() []models.Inversion {
	return s.store.ListarInversiones()
}

func (s *InversionService) BuscarInversionPorID(id int) (models.Inversion, bool) {
	return s.store.BuscarInversionPorID(id)
}

func (s *InversionService) CrearInversion(i models.Inversion) models.Inversion {
	return s.store.CrearInversion(i)
}

func (s *InversionService) ActualizarInversion(id int, datos models.Inversion) (models.Inversion, bool) {
	return s.store.ActualizarInversion(id, datos)
}

func (s *InversionService) BorrarInversion(id int) bool {
	return s.store.BorrarInversion(id)
}
