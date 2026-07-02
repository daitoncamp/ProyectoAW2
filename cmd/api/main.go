package main

import (
	"log"
	"net/http"

	"Proyecto_AWEBII/internal/config"
	"Proyecto_AWEBII/internal/handlers"
	"Proyecto_AWEBII/internal/middleware"
	"Proyecto_AWEBII/internal/models"
	"Proyecto_AWEBII/internal/routes"
	"Proyecto_AWEBII/internal/services"
	"Proyecto_AWEBII/internal/storage"

	"github.com/glebarez/sqlite"
	"github.com/go-chi/chi/v5"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {

	// =====================================================
	// Cargar configuración
	// =====================================================

	cfg := config.Cargar()

	// =====================================================
	// Conexión a la Base de Datos
	// =====================================================

	var (
		db  *gorm.DB
		err error
	)

	switch cfg.DBDriver {

	case "postgres":

		db, err = gorm.Open(postgres.Open(cfg.DBDsn), &gorm.Config{})
		if err != nil {
			log.Fatal(err)
		}

	default:

		db, err = gorm.Open(sqlite.Open(cfg.RutaDB), &gorm.Config{})
		if err != nil {
			log.Fatal(err)
		}

	}

	// =====================================================
	// Crear tablas automáticamente
	// =====================================================

	err = db.AutoMigrate(
		&models.Inversion{},
		&models.TipoInversion{},
		&models.DestinoInversion{},
		&models.Aporte{},
		&models.Usuario{},
	)

	if err != nil {
		log.Fatal("Error en AutoMigrate: ", err)
	}

	// =====================================================
	// Almacén SQLite/PostgreSQL
	// =====================================================

	almacen := storage.NuevoAlmacenSQLite(db)
	almacen.SembrarSiVacio()

	// =====================================================
	// Repositorio de usuarios + Servicio de autenticación
	// =====================================================

	usuarioRepo := storage.NewUsuarioRepository(db)
	authService := services.NuevoAuthService(usuarioRepo)

	// =====================================================
	// Servicios
	// =====================================================

	inversionService := services.NewInversionService(almacen)

	// =====================================================
	// Handlers
	// =====================================================

	authHandler := handlers.NewAuthHandler(authService)
	inversionHandler := handlers.NewInversionHandler(inversionService)

	// =====================================================
	// Router
	// =====================================================

	r := chi.NewRouter()

	r.Use(middleware.Cors)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("API Asociación Estudiantil"))
	})

	r.Get("/test", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Servidor funcionando"))
	})

	// =====================================================
	// Rutas públicas
	// =====================================================

	routes.AuthRoutes(
		r,
		authHandler,
	)

	// =====================================================
	// Rutas protegidas
	// =====================================================

	r.Group(func(r chi.Router) {

		r.Use(middleware.Auth(authService))

		routes.InversionRoutes(
			r,
			inversionHandler,
		)

		// routes.EventoRoutes(...)
		// routes.EstudianteRoutes(...)

	})

	// =====================================================
	// Servidor
	// =====================================================

	log.Println("Servidor ejecutándose en", cfg.Puerto)

	err = http.ListenAndServe(cfg.Puerto, r)

	if err != nil {
		log.Fatal(err)
	}

}
