package routes

import (
	"Proyecto_AWEBII/internal/handlers"

	"github.com/go-chi/chi/v5"
)

func EstudianteRoutes(r chi.Router) {
	r.Get("/estudiantes", handlers.ObtenerEstudiantes)
	r.Post("/estudiantes", handlers.CrearEstudiante)

	r.Get("/estudiantes/{id}", handlers.ObtenerEstudiante)
	r.Put("/estudiantes/{id}", handlers.ActualizarEstudiante)
	r.Delete("/estudiantes/{id}", handlers.EliminarEstudiante)
}
