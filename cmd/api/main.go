package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"

	"Proyecto_AWEBII/internal/routes"
)

func main() {

	r := chi.NewRouter()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("API Asociación Estudiantil"))
	})

	// Ruta de prueba
	r.Get("/test", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Servidor funcionando"))
	})

	// Rutas de estudiantes
	routes.EstudianteRoutes(r)
	routes.EventoRoutes(r)

	log.Println("Servidor ejecutándose en puerto 8080")

	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatal(err)
	}
}
