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
	"Proyecto_AWEBII/internal/storage"
)

func main() {

	// =====================================================
	// Conexión SQLite + GORM
	// =====================================================

	db, err := gorm.Open(sqlite.Open("asociacion.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Error al abrir SQLite: ", err)
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
	// Almacén SQLite
	// =====================================================

	almacen := storage.NuevoAlmacenSQLite(db)
	almacen.SembrarSiVacio()

	// =====================================================
	// Repositorio de usuarios + Servicio de autenticación
	// =====================================================

	usuarioRepo := storage.NewUsuarioRepository(db)
	authService := services.NuevoAuthService(usuarioRepo)

	// Servicios
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

		// Cuando tu compañero termine:
		// routes.EventoRoutes(
		//     r,
		//     handlers.NewEventoHandler(...)
		// )

		// routes.EstudianteRoutes(r)

	})

	// =====================================================
	// Servidor
	// =====================================================

	log.Println("Servidor ejecutándose en puerto 8080")

	err = http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatal(err)
	}
}
