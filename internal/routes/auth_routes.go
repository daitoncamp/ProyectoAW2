package routes

import (
	"Proyecto_AWEBII/internal/handlers"

	"github.com/go-chi/chi/v5"
)

// AuthRoutes registra las rutas relacionadas con autenticación.
func AuthRoutes(r chi.Router, h *handlers.AuthHandler) {

	r.Route("/api/v1/auth", func(r chi.Router) {

		r.Post("/registrar", h.Registrar)
		r.Post("/login", h.Login)

	})
}
