package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"

	"Proyecto_AWEBII/internal/models"
	"Proyecto_AWEBII/internal/services"
)

type EventoHandler struct {
	Service *services.EventoService
}

func NewEventoHandler(s *services.EventoService) *EventoHandler {
	return &EventoHandler{Service: s}
}

// RESPUESTAS
func respondJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(data)
}

func respondError(w http.ResponseWriter, status int, msg string) {
	http.Error(w, msg, status)
}

// LISTAR
func (h *EventoHandler) ListarEventos(w http.ResponseWriter, r *http.Request) {
	eventos, err := h.Service.Listar()
	if err != nil {
		respondError(w, http.StatusInternalServerError, "error al listar eventos")
		return
	}
	respondJSON(w, http.StatusOK, eventos)
}

// OBTENER
func (h *EventoHandler) ObtenerEvento(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		respondError(w, http.StatusBadRequest, "id inválido")
		return
	}

	evento, err := h.Service.ObtenerPorID(id)
	if err != nil {
		if errors.Is(err, services.ErrNoEncontrado) {
			respondError(w, http.StatusNotFound, "evento no encontrado")
			return
		}
		respondError(w, http.StatusInternalServerError, "error interno")
		return
	}

	respondJSON(w, http.StatusOK, evento)
}

// CREAR
func (h *EventoHandler) CrearEvento(w http.ResponseWriter, r *http.Request) {
	var evento models.Evento

	if err := json.NewDecoder(r.Body).Decode(&evento); err != nil {
		respondError(w, http.StatusBadRequest, "JSON inválido")
		return
	}

	creado, err := h.Service.Crear(evento)
	if err != nil {
		if errors.Is(err, services.ErrNombreVacio) {
			respondError(w, http.StatusBadRequest, "nombre vacío")
			return
		}
		respondError(w, http.StatusInternalServerError, "error al crear evento")
		return
	}

	respondJSON(w, http.StatusCreated, creado)
}

// ACTUALIZAR
func (h *EventoHandler) ActualizarEvento(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		respondError(w, http.StatusBadRequest, "id inválido")
		return
	}

	var evento models.Evento
	if err := json.NewDecoder(r.Body).Decode(&evento); err != nil {
		respondError(w, http.StatusBadRequest, "JSON inválido")
		return
	}

	evento.ID = id
	actualizado, err := h.Service.Actualizar(id, evento)
	if err != nil {
		if errors.Is(err, services.ErrNoEncontrado) {
			respondError(w, http.StatusNotFound, "evento no encontrado")
			return
		}
		respondError(w, http.StatusInternalServerError, "error al actualizar")
		return
	}

	respondJSON(w, http.StatusOK, actualizado)
}

// ELIMINAR
func (h *EventoHandler) EliminarEvento(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		respondError(w, http.StatusBadRequest, "id inválido")
		return
	}

	err = h.Service.Eliminar(id)
	if err != nil {
		if errors.Is(err, services.ErrNoEncontrado) {
			respondError(w, http.StatusNotFound, "evento no encontrado")
			return
		}
		respondError(w, http.StatusInternalServerError, "error al eliminar")
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
