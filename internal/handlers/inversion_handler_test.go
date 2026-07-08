package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/go-chi/chi/v5"

	"Proyecto_AWEBII/internal/mocks"
	"Proyecto_AWEBII/internal/models"
	"Proyecto_AWEBII/internal/services"
)

// nuevoHandlerPrueba crea un handler utilizando un almacén simulado (Mock).
// Esto permite probar los endpoints sin necesidad de una base de datos real.
func nuevoHandlerPrueba() *InversionHandler {

	store := &mocks.MockAlmacen{}

	service := services.NewInversionService(store)

	return NewInversionHandler(service)
}

// TestListarInversiones verifica que el endpoint devuelva todas las inversiones.
func TestListarInversiones(t *testing.T) {

	// Creamos un almacén simulado.
	store := &mocks.MockAlmacen{}

	// Insertamos una inversión de prueba.
	store.Inversiones = append(store.Inversiones, models.Inversion{
		ID:                  1,
		Nombre:              "Compra de Equipos",
		MontoInicial:        5000,
		MontoActual:         5200,
		RendimientoEsperado: 8,
		Estado:              "Activa",
		TipoInversionID:     1,
		DestinoInversionID:  1,
	})

	// Creamos el servicio y el handler.
	service := services.NewInversionService(store)
	handler := NewInversionHandler(service)

	// Simulamos una petición GET.
	req := httptest.NewRequest(http.MethodGet, "/inversiones", nil)

	// Capturamos la respuesta.
	rr := httptest.NewRecorder()

	// Ejecutamos el handler.
	handler.ListarInversiones(rr, req)

	// Verificamos que responda HTTP 200.
	if rr.Code != http.StatusOK {
		t.Fatalf("Se esperaba %d pero llegó %d", http.StatusOK, rr.Code)
	}

	// Convertimos la respuesta JSON.
	var inversiones []models.Inversion

	if err := json.NewDecoder(rr.Body).Decode(&inversiones); err != nil {
		t.Fatalf("Error al leer la respuesta JSON: %v", err)
	}

	// Debe existir exactamente una inversión.
	if len(inversiones) != 1 {
		t.Fatalf("Se esperaba 1 inversión pero llegaron %d", len(inversiones))
	}

	// Verificamos el nombre.
	if inversiones[0].Nombre != "Compra de Equipos" {
		t.Fatalf("Nombre incorrecto: %s", inversiones[0].Nombre)
	}
}

// TestBuscarInversionPorID verifica que el endpoint devuelva
// una inversión cuando el ID existe.
afunc TestBuscarInversionPorID(t *testing.T) {

	// Creamos el almacén simulado.
	store := &mocks.MockAlmacen{}

	// Insertamos una inversión de prueba.
	store.Inversiones = append(store.Inversiones, models.Inversion{
		ID:                  1,
		Nombre:              "Proyecto Biblioteca",
		MontoInicial:        8000,
		MontoActual:         8200,
		RendimientoEsperado: 6,
		Estado:              "Activa",
		TipoInversionID:     1,
		DestinoInversionID:  1,
	})

	// Creamos el servicio y el handler.
	service := services.NewInversionService(store)
	handler := NewInversionHandler(service)

	// Simulamos la petición.
	req := httptest.NewRequest(http.MethodGet, "/inversiones/1", nil)

	// Agregamos el parámetro {id} como lo hace Chi.
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("id", "1")

	req = req.WithContext(context.WithValue(
		req.Context(),
		chi.RouteCtxKey,
		rctx,
	))

	rr := httptest.NewRecorder()

	// Ejecutamos el handler.
	handler.BuscarInversionPorID(rr, req)

	// Debe responder HTTP 200.
	if rr.Code != http.StatusOK {
		t.Fatalf("Se esperaba %d pero llegó %d", http.StatusOK, rr.Code)
	}

	// Convertimos el JSON.
	var inversion models.Inversion

	if err := json.NewDecoder(rr.Body).Decode(&inversion); err != nil {
		t.Fatalf("Error al leer JSON: %v", err)
	}

	// Verificamos los datos.
	if inversion.ID != 1 {
		t.Fatalf("Se esperaba ID 1 pero llegó %d", inversion.ID)
	}

	if inversion.Nombre != "Proyecto Biblioteca" {
		t.Fatalf("Nombre incorrecto: %s", inversion.Nombre)
	}
}

// TestCrearInversion verifica que el endpoint cree correctamente
// una nueva inversión cuando el JSON es válido.
func TestCrearInversion(t *testing.T) {

	// Creamos el almacén simulado.
	store := &mocks.MockAlmacen{}

	// Creamos el servicio y el handler.
	service := services.NewInversionService(store)
	handler := NewInversionHandler(service)

	// JSON que simula el cuerpo de una petición POST.
	body := `{
		"nombre":"Nueva Inversión",
		"monto_inicial":5000,
		"monto_actual":5000,
		"rendimiento_esperado":8,
		"estado":"Activa",
		"tipo_inversion_id":1,
		"destino_inversion_id":1
	}`

	// Creamos la petición HTTP.
	req := httptest.NewRequest(
		http.MethodPost,
		"/inversiones",
		strings.NewReader(body),
	)

	req.Header.Set("Content-Type", "application/json")

	// Capturamos la respuesta.
	rr := httptest.NewRecorder()

	// Ejecutamos el handler.
	handler.CrearInversion(rr, req)

	// Debe responder HTTP 201.
	if rr.Code != http.StatusCreated {
		t.Fatalf("Se esperaba %d pero llegó %d", http.StatusCreated, rr.Code)
	}

	// Convertimos el JSON recibido.
	var inversion models.Inversion

	if err := json.NewDecoder(rr.Body).Decode(&inversion); err != nil {
		t.Fatalf("Error al leer JSON: %v", err)
	}

	// Verificamos que tenga ID.
	if inversion.ID == 0 {
		t.Fatal("Se esperaba un ID generado")
	}

	// Verificamos el nombre.
	if inversion.Nombre != "Nueva Inversión" {
		t.Fatalf("Nombre incorrecto: %s", inversion.Nombre)
	}
}

// TestActualizarInversion verifica que el endpoint actualice
// correctamente una inversión existente.
func TestActualizarInversion(t *testing.T) {

	// Creamos el almacén simulado.
	store := &mocks.MockAlmacen{}

	// Insertamos una inversión existente.
	store.Inversiones = append(store.Inversiones, models.Inversion{
		ID:                  1,
		Nombre:              "Inversión Antigua",
		MontoInicial:        1000,
		MontoActual:         1000,
		RendimientoEsperado: 5,
		Estado:              "Activa",
		TipoInversionID:     1,
		DestinoInversionID:  1,
	})

	service := services.NewInversionService(store)
	handler := NewInversionHandler(service)

	// JSON con los nuevos datos.
	body := `{
		"nombre":"Inversión Actualizada",
		"monto_inicial":8000,
		"monto_actual":8200,
		"rendimiento_esperado":10,
		"estado":"Activa",
		"tipo_inversion_id":1,
		"destino_inversion_id":1
	}`

	req := httptest.NewRequest(
		http.MethodPut,
		"/inversiones/1",
		strings.NewReader(body),
	)

	// Simular el parámetro {id} de Chi.
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("id", "1")

	req = req.WithContext(
		context.WithValue(req.Context(), chi.RouteCtxKey, rctx),
	)

	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()

	handler.ActualizarInversion(rr, req)

	if rr.Code != http.StatusOK {
		t.Fatalf("Se esperaba %d pero llegó %d", http.StatusOK, rr.Code)
	}

	var inversion models.Inversion

	if err := json.NewDecoder(rr.Body).Decode(&inversion); err != nil {
		t.Fatalf("Error leyendo JSON: %v", err)
	}

	if inversion.Nombre != "Inversión Actualizada" {
		t.Fatalf("El nombre no se actualizó")
	}

	if inversion.MontoInicial != 8000 {
		t.Fatalf("El monto no se actualizó")
	}
}

// TestBorrarInversion verifica que el endpoint elimine
// correctamente una inversión existente.
func TestBorrarInversion(t *testing.T) {

	// Creamos un almacén simulado.
	store := &mocks.MockAlmacen{}

	// Insertamos una inversión.
	store.Inversiones = append(store.Inversiones, models.Inversion{
		ID:                  1,
		Nombre:              "Inversión",
		MontoInicial:        5000,
		MontoActual:         5200,
		RendimientoEsperado: 8,
		Estado:              "Activa",
		TipoInversionID:     1,
		DestinoInversionID:  1,
	})

	service := services.NewInversionService(store)
	handler := NewInversionHandler(service)

	req := httptest.NewRequest(
		http.MethodDelete,
		"/inversiones/1",
		nil,
	)

	// Simular el parámetro {id} del router Chi.
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("id", "1")

	req = req.WithContext(
		context.WithValue(req.Context(), chi.RouteCtxKey, rctx),
	)

	rr := httptest.NewRecorder()

	handler.BorrarInversion(rr, req)

	if rr.Code != http.StatusOK {
		t.Fatalf("Se esperaba %d pero llegó %d", http.StatusOK, rr.Code)
	}

	if len(store.Inversiones) != 0 {
		t.Fatalf("La inversión no fue eliminada")
	}
}
