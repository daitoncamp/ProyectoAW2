package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"Proyecto_AWEBII/internal/models"
	"Proyecto_AWEBII/internal/storage"

	"github.com/go-chi/chi/v5"
)

type EventoHandler struct {
	store storage.Almacen
}

func NewEventoHandler(store storage.Almacen) *EventoHandler {
	return &EventoHandler{store: store}
}

// LISTAR EVENTOS
func (h *EventoHandler) ObtenerEventos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(h.store.ListarEventos())
}

// CREAR EVENTO
func (h *EventoHandler) CrearEvento(w http.ResponseWriter, r *http.Request) {
	var e models.Evento

	if err := json.NewDecoder(r.Body).Decode(&e); err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}

	creado := h.store.CrearEvento(e)
	json.NewEncoder(w).Encode(creado)
}

// OBTENER POR ID
func (h *EventoHandler) ObtenerEventoPorID(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	ev, ok := h.store.BuscarEventoPorID(id)
	if !ok {
		http.Error(w, "evento no encontrado", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(ev)
}

// ACTUALIZAR EVENTO
func (h *EventoHandler) ActualizarEvento(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	var e models.Evento
	if err := json.NewDecoder(r.Body).Decode(&e); err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}

	actualizado, ok := h.store.ActualizarEvento(id, e)
	if !ok {
		http.Error(w, "evento no encontrado", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(actualizado)
}

// ELIMINAR EVENTO
func (h *EventoHandler) EliminarEvento(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	ok := h.store.BorrarEvento(id)
	if !ok {
		http.Error(w, "evento no encontrado", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
