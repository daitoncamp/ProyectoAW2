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

func (s *InversionService) BuscarInversionPorID(id int) (models.Inversion, error) {

	inversion, ok := s.store.BuscarInversionPorID(id)

	if !ok {
		return models.Inversion{}, ErrNoEncontrado
	}

	return inversion, nil
}

// ==========================================
// Crear inversión
// ==========================================

func (s *InversionService) CrearInversion(i models.Inversion) (models.Inversion, error) {

	// Validaciones de negocio
	if i.Nombre == "" {
		return models.Inversion{}, ErrNombreVacio
	}

	if i.MontoInicial <= 0 {
		return models.Inversion{}, ErrMontoInvalido
	}

	if i.Estado == "" {
		return models.Inversion{}, ErrEstadoVacio
	}

	// Guardar
	return s.store.CrearInversion(i), nil
}

func (s *InversionService) ActualizarInversion(id int, datos models.Inversion) (models.Inversion, error) {

	if datos.Nombre == "" {
		return models.Inversion{}, ErrNombreVacio
	}

	if datos.MontoInicial <= 0 {
		return models.Inversion{}, ErrMontoInvalido
	}

	if datos.Estado == "" {
		return models.Inversion{}, ErrEstadoVacio
	}

	inversion, ok := s.store.ActualizarInversion(id, datos)
	if !ok {
		return models.Inversion{}, ErrNoEncontrado
	}

	return inversion, nil
}

func (s *InversionService) BorrarInversion(id int) error {

	ok := s.store.BorrarInversion(id)

	if !ok {
		return ErrNoEncontrado
	}

	return nil
}
