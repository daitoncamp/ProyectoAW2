package routes

import (
	"Proyecto_AWEBII/internal/handlers"

	"github.com/go-chi/chi/v5"
)

func InversionRoutes(r chi.Router, h *handlers.InversionHandler) {

	r.Route("/api/v1/inversiones", func(r chi.Router) {
		r.Get("/", h.ListarInversiones)
		r.Post("/", h.CrearInversion)
		r.Get("/{id}", h.BuscarInversionPorID)
		r.Put("/{id}", h.ActualizarInversion)
		r.Delete("/{id}", h.BorrarInversion)
	})
}
