package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"

	"Proyecto_AWEBII/internal/handlers"
	"Proyecto_AWEBII/internal/service"
	"Proyecto_AWEBII/internal/storage"
)

func main() {

	repo := storage.NewMemoria()

	eventoService := service.NewEventoService(repo)
	inversionService := service.NewInversionService(repo)

	eventoHandler := handlers.NewEventoHandler(eventoService)
	inversionHandler := handlers.NewInversionHandler(inversionService)

	r := chi.NewRouter()

	r.Route("/api/v1", func(r chi.Router) {

		// EVENTOS
		r.Get("/eventos", eventoHandler.ListarEventos)
		r.Get("/eventos/{id}", eventoHandler.ObtenerEvento)
		r.Post("/eventos", eventoHandler.CrearEvento)
		r.Put("/eventos/{id}", eventoHandler.ActualizarEvento)
		r.Delete("/eventos/{id}", eventoHandler.EliminarEvento)

		// INVERSIONES
		r.Get("/inversiones", inversionHandler.ListarInversiones)
		r.Get("/inversiones/{id}", inversionHandler.ObtenerInversion)
		r.Post("/inversiones", inversionHandler.CrearInversion)
		r.Put("/inversiones/{id}", inversionHandler.ActualizarInversion)
		r.Delete("/inversiones/{id}", inversionHandler.EliminarInversion)
	})

	http.ListenAndServe(":8080", r)
}
