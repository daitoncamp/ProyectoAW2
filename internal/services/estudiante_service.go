package services

import (
	"errors"

	"Proyecto_AWEBII/internal/models"
	"Proyecto_AWEBII/internal/storage"
)

// OBTENER TODOS LOS ESTUDIANTES
func ObtenerEstudiantes() []models.Estudiante {
	return storage.Estudiantes
}

// CREAR ESTUDIANTE
func CrearEstudiante(estudiante models.Estudiante) {

	// Generar ID automáticamente
	if len(storage.Estudiantes) == 0 {
		estudiante.ID = 1
	} else {
		estudiante.ID = storage.Estudiantes[len(storage.Estudiantes)-1].ID + 1
	}

	storage.Estudiantes = append(storage.Estudiantes, estudiante)
}

// OBTENER ESTUDIANTE POR ID
func ObtenerEstudiantePorID(id int) (models.Estudiante, error) {

	for _, estudiante := range storage.Estudiantes {

		if estudiante.ID == id {
			return estudiante, nil
		}
	}

	return models.Estudiante{}, errors.New("estudiante no encontrado")
}

// ACTUALIZAR ESTUDIANTE
func ActualizarEstudiante(id int, estudianteActualizado models.Estudiante) error {

	for i, estudiante := range storage.Estudiantes {

		if estudiante.ID == id {

			// Mantener el mismo ID
			estudianteActualizado.ID = id

			storage.Estudiantes[i] = estudianteActualizado

			return nil
		}
	}

	return errors.New("estudiante no encontrado")
}

// ELIMINAR ESTUDIANTE
func EliminarEstudiante(id int) error {

	for i, estudiante := range storage.Estudiantes {

		if estudiante.ID == id {

			storage.Estudiantes = append(
				storage.Estudiantes[:i],
				storage.Estudiantes[i+1:]...,
			)

			return nil
		}
	}

	return errors.New("estudiante no encontrado")
}