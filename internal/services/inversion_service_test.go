package services

import (
	"testing"

	"Proyecto_AWEBII/internal/mocks"
	"Proyecto_AWEBII/internal/models"
)

// =====================================================
// CREAR INVERSIÓN
// =====================================================

// Verifica que una inversión válida se cree correctamente.
func TestCrearInversion(t *testing.T) {

	// Se crea un almacén simulado (Mock).
	store := &mocks.MockAlmacen{}

	// Se crea el servicio utilizando el Mock.
	service := NewInversionService(store)

	// Datos válidos para crear la inversión.
	inversion := models.Inversion{
		Nombre:              "Compra de Computadoras",
		MontoInicial:        5000,
		MontoActual:         5200,
		RendimientoEsperado: 8,
		Estado:              "Activa",
		TipoInversionID:     1,
		DestinoInversionID:  1,
	}

	// Se ejecuta el método que se desea probar.
	nueva, err := service.CrearInversion(inversion)

	// No debe existir ningún error.
	if err != nil {
		t.Fatalf("No esperaba error: %v", err)
	}

	// El Mock debe asignar un ID automáticamente.
	if nueva.ID == 0 {
		t.Error("Se esperaba un ID generado")
	}

	// Se verifica que el nombre se haya almacenado correctamente.
	if nueva.Nombre != inversion.Nombre {
		t.Error("El nombre no coincide")
	}
}

// Verifica que no se permita crear una inversión sin nombre.
func TestCrearInversionNombreVacio(t *testing.T) {

	store := &mocks.MockAlmacen{}
	service := NewInversionService(store)

	inversion := models.Inversion{
		Nombre:              "",
		MontoInicial:        1000,
		MontoActual:         1000,
		RendimientoEsperado: 5,
		Estado:              "Activa",
		TipoInversionID:     1,
		DestinoInversionID:  1,
	}

	_, err := service.CrearInversion(inversion)

	// Debe producirse un error.
	if err == nil {
		t.Fatal("Se esperaba un error por nombre vacío")
	}

	// El error debe ser exactamente el esperado.
	if err != ErrNombreVacio {
		t.Fatalf("Se esperaba %v pero llegó %v", ErrNombreVacio, err)
	}
}

// Verifica que el monto inicial sea mayor que cero.
func TestCrearInversionMontoInvalido(t *testing.T) {

	store := &mocks.MockAlmacen{}
	service := NewInversionService(store)

	inversion := models.Inversion{
		Nombre:              "Inversión de prueba",
		MontoInicial:        0,
		MontoActual:         0,
		RendimientoEsperado: 5,
		Estado:              "Activa",
		TipoInversionID:     1,
		DestinoInversionID:  1,
	}

	_, err := service.CrearInversion(inversion)

	if err == nil {
		t.Fatal("Se esperaba un error por monto inválido")
	}

	if err != ErrMontoInvalido {
		t.Fatalf("Se esperaba %v pero llegó %v", ErrMontoInvalido, err)
	}
}

// Verifica que el estado sea obligatorio.
func TestCrearInversionEstadoVacio(t *testing.T) {

	store := &mocks.MockAlmacen{}
	service := NewInversionService(store)

	inversion := models.Inversion{
		Nombre:              "Inversión de prueba",
		MontoInicial:        1000,
		MontoActual:         1000,
		RendimientoEsperado: 5,
		Estado:              "",
		TipoInversionID:     1,
		DestinoInversionID:  1,
	}

	_, err := service.CrearInversion(inversion)

	if err == nil {
		t.Fatal("Se esperaba un error por estado vacío")
	}

	if err != ErrEstadoVacio {
		t.Fatalf("Se esperaba %v pero llegó %v", ErrEstadoVacio, err)
	}
}

// =====================================================
// BUSCAR INVERSIÓN
// =====================================================

// Verifica que una inversión existente pueda recuperarse por su ID.
func TestBuscarInversionPorID(t *testing.T) {

	store := &mocks.MockAlmacen{}

	// Se agrega una inversión al Mock.
	store.Inversiones = append(store.Inversiones, models.Inversion{
		ID:                  1,
		Nombre:              "Fondo Universitario",
		MontoInicial:        1000,
		MontoActual:         1200,
		RendimientoEsperado: 10,
		Estado:              "Activa",
		TipoInversionID:     1,
		DestinoInversionID:  1,
	})

	service := NewInversionService(store)

	inversion, err := service.BuscarInversionPorID(1)

	if err != nil {
		t.Fatalf("No se esperaba error: %v", err)
	}

	if inversion.Nombre != "Fondo Universitario" {
		t.Fatalf("Se esperaba 'Fondo Universitario' pero llegó '%s'", inversion.Nombre)
	}
}

// Verifica que se devuelva un error cuando la inversión no existe.
func TestBuscarInversionNoExiste(t *testing.T) {

	store := &mocks.MockAlmacen{}
	service := NewInversionService(store)

	_, err := service.BuscarInversionPorID(99)

	if err == nil {
		t.Fatal("Se esperaba un error")
	}

	if err != ErrNoEncontrado {
		t.Fatalf("Se esperaba %v pero llegó %v", ErrNoEncontrado, err)
	}
}

// =====================================================
// ACTUALIZAR INVERSIÓN
// =====================================================

// Verifica que una inversión existente pueda actualizarse correctamente.
func TestActualizarInversion(t *testing.T) {

	store := &mocks.MockAlmacen{}

	// Se registra una inversión inicial.
	store.Inversiones = append(store.Inversiones, models.Inversion{
		ID:                  1,
		Nombre:              "Inversión vieja",
		MontoInicial:        1000,
		MontoActual:         1000,
		RendimientoEsperado: 10,
		Estado:              "Activa",
		TipoInversionID:     1,
		DestinoInversionID:  1,
	})

	service := NewInversionService(store)

	// Nuevos datos para actualizar.
	datos := models.Inversion{
		Nombre:              "Inversión actualizada",
		MontoInicial:        2000,
		MontoActual:         2100,
		RendimientoEsperado: 12,
		Estado:              "Activa",
		TipoInversionID:     1,
		DestinoInversionID:  1,
	}

	inversion, err := service.ActualizarInversion(1, datos)

	if err != nil {
		t.Fatalf("No se esperaba error: %v", err)
	}

	if inversion.Nombre != "Inversión actualizada" {
		t.Fatalf("Se esperaba 'Inversión actualizada' pero llegó '%s'", inversion.Nombre)
	}

	if inversion.MontoInicial != 2000 {
		t.Fatalf("Se esperaba monto 2000 pero llegó %.2f", inversion.MontoInicial)
	}
}

// Verifica que no sea posible actualizar una inversión inexistente.
func TestActualizarInversionNoExiste(t *testing.T) {

	store := &mocks.MockAlmacen{}
	service := NewInversionService(store)

	datos := models.Inversion{
		Nombre:              "Nueva",
		MontoInicial:        1000,
		MontoActual:         1000,
		RendimientoEsperado: 5,
		Estado:              "Activa",
		TipoInversionID:     1,
		DestinoInversionID:  1,
	}

	_, err := service.ActualizarInversion(50, datos)

	if err == nil {
		t.Fatal("Se esperaba un error")
	}

	if err != ErrNoEncontrado {
		t.Fatalf("Se esperaba %v pero llegó %v", ErrNoEncontrado, err)
	}
}

// =====================================================
// ELIMINAR INVERSIÓN
// =====================================================

// Verifica que una inversión existente pueda eliminarse correctamente.
func TestBorrarInversion(t *testing.T) {

	store := &mocks.MockAlmacen{}

	store.Inversiones = append(store.Inversiones, models.Inversion{
		ID:                  1,
		Nombre:              "Fondo",
		MontoInicial:        1000,
		MontoActual:         1100,
		RendimientoEsperado: 10,
		Estado:              "Activa",
		TipoInversionID:     1,
		DestinoInversionID:  1,
	})

	service := NewInversionService(store)

	err := service.BorrarInversion(1)

	if err != nil {
		t.Fatalf("No se esperaba error: %v", err)
	}

	// Después de eliminar, el Mock debe quedar vacío.
	if len(store.Inversiones) != 0 {
		t.Fatalf("Se esperaba que no existan inversiones")
	}
}

// Verifica que se retorne un error al intentar eliminar una inversión inexistente.
func TestBorrarInversionNoExiste(t *testing.T) {

	store := &mocks.MockAlmacen{}
	service := NewInversionService(store)

	err := service.BorrarInversion(50)

	if err == nil {
		t.Fatal("Se esperaba un error")
	}

	if err != ErrNoEncontrado {
		t.Fatalf("Se esperaba %v pero llegó %v", ErrNoEncontrado, err)
	}
}
