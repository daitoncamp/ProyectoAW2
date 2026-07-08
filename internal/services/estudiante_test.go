package services

import (
	"testing"

	"Proyecto_AWEBII/internal/storage"
)

func TestObtenerEstudiantePorID(t *testing.T) {
	// Verificamos que existan datos de prueba
	if len(storage.Estudiantes) == 0 {
		t.Fatal("no existen estudiantes de prueba")
	}

	id := storage.Estudiantes[0].ID

	estudiante, err := ObtenerEstudiantePorID(id)

	if err != nil {
		t.Fatalf("se esperaba encontrar el estudiante y ocurrió un error: %v", err)
	}

	if estudiante.ID != 999 {
		t.Errorf("se esperaba ID %d, se obtuvo %d", id, estudiante.ID)
	}
}
