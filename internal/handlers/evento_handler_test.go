package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"Proyecto_AWEBII/internal/mocks"
	"Proyecto_AWEBII/internal/models"
	"Proyecto_AWEBII/internal/services"

	"github.com/go-chi/chi/v5"
)

func TestCrearEventoHandler(t *testing.T) {

	mock := &mocks.MockEventoAlmacen{}

	service := services.NewEventoService(mock)

	handler := NewEventoHandler(service)

	router := chi.NewRouter()

	router.Post("/eventos", handler.CrearEvento)

	body := models.Evento{
		Nombre: "Conferencia Go",
		Fecha:  "2026-07-15",
		Lugar:  "Auditorio",
	}

	jsonBody, _ := json.Marshal(body)

	req := httptest.NewRequest(
		http.MethodPost,
		"/eventos",
		bytes.NewBuffer(jsonBody),
	)

	req.Header.Set(
		"Content-Type",
		"application/json",
	)

	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf(
			"se esperaba 200 pero llegó %d",
			rec.Code,
		)
	}

	if len(mock.Eventos) != 1 {
		t.Fatalf(
			"se esperaba guardar un evento",
		)
	}
}

// segundo test listar
func TestListarEventosHandler(t *testing.T) {

	mock := &mocks.MockEventoAlmacen{}

	// agregamos datos al mock antes de llamar al handler
	mock.Eventos = append(mock.Eventos, models.Evento{
		ID:     1,
		Nombre: "Conferencia Go",
		Fecha:  "2026-07-15",
		Lugar:  "Auditorio",
	})

	service := services.NewEventoService(mock)
	handler := NewEventoHandler(service)

	router := chi.NewRouter()
	router.Get("/eventos", handler.ListarEventos)

	req := httptest.NewRequest(
		http.MethodGet,
		"/eventos",
		nil,
	)

	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf(
			"se esperaba 200 pero llegó %d",
			rec.Code,
		)
	}

	//asta aqui

	var respuesta []models.Evento
	err := json.NewDecoder(rec.Body).Decode(&respuesta)
	if err != nil {
		t.Fatalf("error decodificando respuesta")
	}

	if len(respuesta) != 1 {
		t.Fatalf("se esperaba 1 evento")
	}
}

// tercer test obenter
func TestObtenerEventoHandler(t *testing.T) {

	mock := &mocks.MockEventoAlmacen{}

	// precargamos un evento
	mock.Eventos = []models.Evento{
		{
			ID:     1,
			Nombre: "Conferencia Go",
			Fecha:  "2026-07-15",
			Lugar:  "Auditorio",
		},
	}

	service := services.NewEventoService(mock)
	handler := NewEventoHandler(service)

	router := chi.NewRouter()
	router.Get("/eventos/{id}", handler.ObtenerEvento)

	req := httptest.NewRequest(http.MethodGet, "/eventos/1", nil)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("se esperaba 200 pero llegó %d", rec.Code)
	}

	var evento models.Evento
	err := json.NewDecoder(rec.Body).Decode(&evento)
	if err != nil {
		t.Fatalf("error decodificando respuesta")
	}

	if evento.ID != 1 {
		t.Fatalf("se esperaba ID 1 pero llegó %d", evento.ID)
	}
}

// test de actualizar, cuarto
func TestActualizarEventoHandler(t *testing.T) {

	mock := &mocks.MockEventoAlmacen{}

	// evento inicial
	mock.Eventos = []models.Evento{
		{
			ID:     1,
			Nombre: "Viejo evento",
			Fecha:  "2026-07-10",
			Lugar:  "Auditorio",
		},
	}

	service := services.NewEventoService(mock)
	handler := NewEventoHandler(service)

	router := chi.NewRouter()
	router.Put("/eventos/{id}", handler.ActualizarEvento)

	body := models.Evento{
		Nombre: "Evento actualizado",
		Fecha:  "2026-07-20",
		Lugar:  "Sala principal",
	}

	jsonBody, _ := json.Marshal(body)

	req := httptest.NewRequest(http.MethodPut, "/eventos/1", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("se esperaba 200 pero llegó %d", rec.Code)
	}

	var actualizado models.Evento
	err := json.NewDecoder(rec.Body).Decode(&actualizado)
	if err != nil {
		t.Fatalf("error decodificando respuesta")
	}

	if actualizado.Nombre != "Evento actualizado" {
		t.Fatalf("no se actualizó correctamente")
	}
}

// quinto test: eliminar
func TestEliminarEventoHandler(t *testing.T) {

	mock := &mocks.MockEventoAlmacen{}

	// precargamos un evento
	mock.Eventos = []models.Evento{
		{
			ID:     1,
			Nombre: "Evento a eliminar",
			Fecha:  "2026-07-10",
			Lugar:  "Auditorio",
		},
	}

	service := services.NewEventoService(mock)
	handler := NewEventoHandler(service)

	router := chi.NewRouter()
	router.Delete("/eventos/{id}", handler.EliminarEvento)

	req := httptest.NewRequest(http.MethodDelete, "/eventos/1", nil)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusNoContent {
		t.Fatalf("se esperaba 204 pero llegó %d", rec.Code)
	}

	// validación extra: ya no debe existir el evento
	if len(mock.Eventos) != 0 {
		t.Fatalf("se esperaba que el evento se elimine")
	}
}
