package main

import (
	"log"
	"net/http"

	"Proyecto_AWEBII/internal/handlers"
	"Proyecto_AWEBII/internal/models"
	"Proyecto_AWEBII/internal/services"
	"Proyecto_AWEBII/internal/storage"

	"github.com/glebarez/sqlite"
	"github.com/go-chi/chi/v5"
	chimw "github.com/go-chi/chi/v5/middleware"
	"gorm.io/gorm"
)

func main() {
	gdb, err := gorm.Open(sqlite.Open("asociacion.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("no se pudo abrir la base de datos: ", err)
	}

	if err := gdb.AutoMigrate(
		&models.Evento{},
		&models.CategoriaEvento{},
		&models.Asistencia{},
		&models.Usuario{},
	); err != nil {
		log.Fatal("falló AutoMigrate: ", err)
	}

	almacen := storage.NuevoAlmacenSQLite(gdb)

	eventoService := services.NewEventoService(almacen)
	eventoHandler := handlers.NewEventoHandler(eventoService)

	r := chi.NewRouter()

	r.Use(chimw.Logger)
	r.Use(chimw.Recoverer)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("API Asociación Estudiantil funcionando correctamente"))
	})

	log.Println("Backend de almacenamiento: GORM")

	r.Route("/api/v1/eventos", func(r chi.Router) {
		r.Get("/", eventoHandler.ListarEventos)
		r.Post("/", eventoHandler.CrearEvento)
		r.Get("/{id}", eventoHandler.ObtenerEvento)
		r.Put("/{id}", eventoHandler.ActualizarEvento)
		r.Delete("/{id}", eventoHandler.EliminarEvento)
	})

	const addr = ":8080"

	log.Printf("API escuchando en http://localhost%s", addr)

	if err := http.ListenAndServe(addr, r); err != nil {
		log.Fatal(err)
	}
}
