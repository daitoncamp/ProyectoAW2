package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"Proyecto_AWEBII/internal/models"
	"Proyecto_AWEBII/internal/services"

	"github.com/go-chi/chi/v5"
)

type InversionHandler struct {
	service *services.InversionService
}

func NewInversionHandler(service *services.InversionService) *InversionHandler {
	return &InversionHandler{
		service: service,
	}
}

// =====================================================
// Crear inversión
// =====================================================

func (h *InversionHandler) CrearInversion(w http.ResponseWriter, r *http.Request) {

	var inversion models.Inversion

	if err := json.NewDecoder(r.Body).Decode(&inversion); err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}

	nueva, err := h.service.CrearInversion(inversion)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(nueva)
}

// =====================================================
// Listar inversiones
// =====================================================

func (h *InversionHandler) ListarInversiones(w http.ResponseWriter, r *http.Request) {

	inversiones := h.service.ListarInversiones()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(inversiones)
}

// =====================================================
// Buscar inversión por ID
// =====================================================

func (h *InversionHandler) BuscarInversionPorID(w http.ResponseWriter, r *http.Request) {

	idStr := chi.URLParam(r, "id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	inversion, err := h.service.BuscarInversionPorID(id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(inversion)
}

// =====================================================
// Actualizar inversión
// =====================================================

func (h *InversionHandler) ActualizarInversion(w http.ResponseWriter, r *http.Request) {

	idStr := chi.URLParam(r, "id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	var datos models.Inversion

	if err := json.NewDecoder(r.Body).Decode(&datos); err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}

	inversionActualizada, err := h.service.ActualizarInversion(id, datos)

	if err != nil {
		if err == services.ErrNoEncontrado {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(inversionActualizada)
}

// =====================================================
// Eliminar inversión
// =====================================================

func (h *InversionHandler) BorrarInversion(w http.ResponseWriter, r *http.Request) {

	idStr := chi.URLParam(r, "id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	if err := h.service.BorrarInversion(id); err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(map[string]string{
		"mensaje": "Inversión eliminada correctamente",
	})
}
