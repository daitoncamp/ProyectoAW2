package routes

import (
	"Proyecto_AWEBII/internal/handlers"

	"github.com/go-chi/chi/v5"
)

func EstudianteRoutes(r chi.Router) {

	r.Route("/estudiantes", func(r chi.Router) {

		// Obtener todos los estudiantes
		r.Get("/", handlers.ObtenerEstudiantes)

		// Crear un estudiante
		r.Post("/", handlers.CrearEstudiante)

		// Obtener un estudiante por ID
		r.Get("/{id}", handlers.ObtenerEstudiante)

		// Actualizar un estudiante
		r.Put("/{id}", handlers.ActualizarEstudiante)

		// Eliminar un estudiante
		r.Delete("/{id}", handlers.EliminarEstudiante)
	})
}