package storage

import (
	"testing"

	"Proyecto_AWEBII/internal/models"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

// TestRepositorioSQLite verifica que una inversión creada
// realmente se almacene y pueda recuperarse desde SQLite en memoria.
func TestRepositorioSQLite(t *testing.T) {

	// Base de datos temporal en memoria.
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("No se pudo abrir la base de datos: %v", err)
	}

	// Crear la tabla.
	err = db.AutoMigrate(&models.Inversion{})
	if err != nil {
		t.Fatalf("Error en AutoMigrate: %v", err)
	}

	// Repositorio real.
	store := NuevoAlmacenSQLite(db)

	// Crear una inversión.
	inversion := models.Inversion{
		Nombre:              "Inversión Test",
		MontoInicial:        1000,
		MontoActual:         1100,
		RendimientoEsperado: 10,
		Estado:              "Activa",
		TipoInversionID:     1,
		DestinoInversionID:  1,
	}

	nueva := store.CrearInversion(inversion)

	// Buscarla nuevamente.
	encontrada, ok := store.BuscarInversionPorID(nueva.ID)

	if !ok {
		t.Fatal("Se esperaba encontrar la inversión")
	}

	if encontrada.Nombre != inversion.Nombre {
		t.Fatalf("Se esperaba %s pero llegó %s",
			inversion.Nombre,
			encontrada.Nombre,
		)
	}
}
