package services

import (
	"testing"

	"Proyecto_AWEBII/internal/mocks"
	"Proyecto_AWEBII/internal/models"
)

// Test de creación
func TestCrearEventoNombreVacio(t *testing.T) {

	store := &mocks.MockEventoAlmacen{}
	service := NewEventoService(store)

	evento := models.Evento{
		Nombre: "",
		Fecha:  "2026-07-10",
		Lugar:  "Auditorio",
	}

	_, err := service.Crear(evento)

	if err == nil {
		t.Fatal("Se esperaba error por nombre vacío")
	}

	if err != ErrEventoNombreVacio {
		t.Fatalf(
			"Se esperaba %v pero llegó %v",
			ErrEventoNombreVacio,
			err,
		)
	}
}
