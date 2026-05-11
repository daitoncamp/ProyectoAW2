package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {

	// Crear router
	r := chi.NewRouter()

	// Ruta de prueba
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("API Asociación Estudiantil"))
	})

	// Iniciar servidor
	log.Println("Servidor ejecutándose en puerto 8080")

	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatal(err)
	}
}
