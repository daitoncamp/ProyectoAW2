// registrar rutas
package routes

import (
	"Proyecto_AWEBII/internal/handlers"

	"github.com/go-chi/chi/v5"
)

func EventoRoutes(r chi.Router, h *handlers.EventoHandler) {

	r.Route("/api/v1/events", func(r chi.Router) {
		r.Get("/", h.ObtenerEventos)
		r.Post("/", h.CrearEvento)
		r.Get("/{id}", h.ObtenerEventoPorID)
		r.Put("/{id}", h.ActualizarEvento)
		r.Delete("/{id}", h.EliminarEvento)
	})
}
