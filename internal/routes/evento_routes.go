package routes

//registrar rutas

import (
	"Proyecto_AWEBII/internal/handlers"

	"github.com/go-chi/chi/v5"
)

func EventoRoutes(r chi.Router) {
	r.Get("/events", handlers.ObtenerEventos)
	r.Post("/events", handlers.CrearEvento)
	r.Get("/events/{id}", handlers.ObtenerEvento)
	r.Put("/events/{id}", handlers.ActualizarEvento)
	r.Delete("/events/{id}", handlers.EliminarEvento)
}
