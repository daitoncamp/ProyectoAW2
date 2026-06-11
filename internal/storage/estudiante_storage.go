package storage

import (
	"Proyecto_AWEBII/internal/models"
	"time"
)

var Estudiantes = []models.Estudiante{
	{
		ID:              1,
		Nombres:         "Daiton",
		Apellidos:       "Campuzano",
		Email:           "daiton@uleam.com",
		Telefono:        "0999999999",
		FechaNacimiento: time.Date(2003, 5, 10, 0, 0, 0, 0, time.UTC),
		CarreraID:       1,
		FechaRegistro:   time.Now(),
	},
	{
		ID:              2,
		Nombres:         "Julissa",
		Apellidos:       "Armas",
		Email:           "julissa@uleam.com",
		Telefono:        "0988888888",
		FechaNacimiento: time.Date(2002, 8, 15, 0, 0, 0, 0, time.UTC),
		CarreraID:       2,
		FechaRegistro:   time.Now(),
	},
}

var Carreras []models.Carrera
var EstadosAcademicos []models.EstadoAcademico
