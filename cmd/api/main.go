package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"

	"Proyecto_AWEBII/internal/handlers"
	"Proyecto_AWEBII/internal/middleware"
	"Proyecto_AWEBII/internal/routes"
	"Proyecto_AWEBII/internal/storage"
)

func main() {

	r := chi.NewRouter()

	r.Use(middleware.Cors)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("API Asociación Estudiantil"))
	})

	// Memoria
	memoria := storage.NuevaMemoria()
	memoria.SeedInversiones()
	memoria.SeedEventos()

	// Ruta de prueba (opcional de equipo)
	r.Get("/test", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Servidor funcionando"))
	})

	// Rutas del sistema
	routes.EventoRoutes(
		r,
		handlers.NewEventoHandler(memoria),
	)

	routes.InversionRoutes(
		r,
		handlers.NewInversionHandler(memoria),
	)

	// Rutas de estudiantes
	routes.EstudianteRoutes(r)

	log.Println("Servidor ejecutándose en puerto 8080")

	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatal(err)
	}
}
