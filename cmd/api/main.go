package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"

	"Proyecto_AWEBII/internal/routes"
)

func main() {

	// Crear routerr
	r := chi.NewRouter()
	// Ruta de prueba
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("API Asociación Estudiantil"))
	})

	// Rutas del módulo eventos
	routes.EventoRoutes(r)

	// Iniciar servidor
	log.Println("Servidor ejecutándose en puerto 8080")

	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatal(err)
	}
}
