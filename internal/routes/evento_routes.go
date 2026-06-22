package routes

import (
	"Proyecto_AWEBII/internal/handlers"

	"github.com/go-chi/chi/v5"
)

func EventoRoutes(r chi.Router, h *handlers.EventoHandler) {
	r.Route("/api/v1/eventos", func(r chi.Router) {

		r.Get("/", h.ListarEventos)
		r.Post("/", h.CrearEvento)
		r.Get("/{id}", h.ObtenerEvento)
		r.Put("/{id}", h.ActualizarEvento)
		r.Delete("/{id}", h.EliminarEvento)

	})
}
