package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"

	"Proyecto_AWEBII/internal/routes"

	"Proyecto_AWEBII/internal/handlers"
	"Proyecto_AWEBII/internal/storage"
)

func main() {

	r := chi.NewRouter()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("API Asociación Estudiantil"))
	})

	// Memoria de inversiones
	memoria := storage.NuevaMemoria()
	memoria.SeedInversiones()

	// Rutas
	routes.EventoRoutes(r)
	routes.InversionRoutes(
		r,
		handlers.NewInversionHandler(memoria),
	)

	log.Println("Servidor ejecutándose en puerto 8080")

	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatal(err)
	}
}
