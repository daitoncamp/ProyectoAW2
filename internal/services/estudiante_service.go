package services

import (
	"Proyecto_AWEBII/internal/models"
	"Proyecto_AWEBII/internal/storage"
	"errors"
)

func ObtenerEstudiantes() []models.Estudiante {
	return storage.Estudiantes
}

func CrearEstudiante(estudiante models.Estudiante) {
	storage.Estudiantes = append(storage.Estudiantes, estudiante)
}

func ObtenerEstudiantePorID(id int) (models.Estudiante, error) {

	for _, estudiante := range storage.Estudiantes {

		if estudiante.ID == id {
			return estudiante, nil
		}
	}

	return models.Estudiante{}, errors.New("estudiante no encontrado")
}

func ActualizarEstudiante(id int, estudianteActualizado models.Estudiante) error {

	for i, estudiante := range storage.Estudiantes {

		if estudiante.ID == id {
			storage.Estudiantes[i] = estudianteActualizado
			return nil
		}
	}

	return errors.New("estudiante no encontrado")
}

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
