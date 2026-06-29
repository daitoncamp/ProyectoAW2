package main

import (
	"log"
	"net/http"

	"github.com/glebarez/sqlite"
	"github.com/go-chi/chi/v5"
	"gorm.io/gorm"

	"Proyecto_AWEBII/internal/handlers"
	"Proyecto_AWEBII/internal/middleware"
	"Proyecto_AWEBII/internal/models"
	"Proyecto_AWEBII/internal/routes"
	"Proyecto_AWEBII/internal/services"

	//"Proyecto_AWEBII/internal/services"//
	"Proyecto_AWEBII/internal/storage"
)

func main() {

	// Conexión SQLite + GORM

	db, err := gorm.Open(sqlite.Open("asociacion.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Error al abrir SQLite:", err)
	}

	// AutoMigrate

	err = db.AutoMigrate(
		&models.Inversion{},
		&models.TipoInversion{},
		&models.DestinoInversion{},
		&models.Aporte{},

		&models.Evento{},
		&models.CategoriaEvento{},
		&models.Asistencia{},
		&models.Usuario{},
	)
	if err != nil {
		log.Fatal("Error en AutoMigrate:", err)
	}

	// Almacenamiento SQLite

	almacen := storage.NuevoAlmacenSQLite(db)

	// Opcional: insertar datos semilla
	almacen.SembrarSiVacio()

	// SErvicio y el  handler eventos

	eventoService := services.NewEventoService(almacen)
	eventoHandler := handlers.NewEventoHandler(eventoService)

	// Router

	r := chi.NewRouter()

	r.Use(middleware.Cors)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("API Asociación Estudiantil"))
	})

	r.Get("/test", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Servidor funcionando"))
	})

	routes.InversionRoutes(
		r,
		handlers.NewInversionHandler(almacen),
	)

	routes.EventoRoutes(r, eventoHandler)

	log.Println("Servidor ejecutándose en puerto 8080")

	err = http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatal(err)
	}
}
