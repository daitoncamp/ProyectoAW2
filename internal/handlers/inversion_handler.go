package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"Proyecto_AWEBII/internal/models"
	"Proyecto_AWEBII/internal/storage"

	"github.com/go-chi/chi/v5"
)

type InversionHandler struct {
	store storage.Almacen
}

func NewInversionHandler(store storage.Almacen) *InversionHandler {
	return &InversionHandler{
		store: store,
	}
}

// CrearInversion maneja la creación de una nueva inversión a través de una solicitud POST.
func (h *InversionHandler) CrearInversion(w http.ResponseWriter, r *http.Request) {

	var inversion models.Inversion

	err := json.NewDecoder(r.Body).Decode(&inversion)
	if err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}

	if inversion.Nombre == "" {
		http.Error(w, "El nombre es obligatorio", http.StatusBadRequest)
		return
	}

	if inversion.MontoInicial <= 0 {
		http.Error(w, "El monto inicial debe ser mayor a 0", http.StatusBadRequest)
		return
	}

	if inversion.Estado == "" {
		http.Error(w, "El estado es obligatorio", http.StatusBadRequest)
		return
	}

	nueva := h.store.CrearInversion(inversion)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(nueva)
}

// ListarInversiones maneja la solicitud GET para listar todas las inversiones.
func (h *InversionHandler) ListarInversiones(w http.ResponseWriter, r *http.Request) {

	inversiones := h.store.ListarInversiones()

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(inversiones)
}

// BuscarInversionPorID maneja la solicitud GET para obtener una inversión específica por su ID.
func (h *InversionHandler) BuscarInversionPorID(w http.ResponseWriter, r *http.Request) {

	idStr := chi.URLParam(r, "id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	inversion, encontrado := h.store.BuscarInversionPorID(id)

	if !encontrado {
		http.Error(w, "Inversión no encontrada", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(inversion)
}

// ActualizarInversion maneja la solicitud PUT para actualizar una inversión existente por su ID.
func (h *InversionHandler) ActualizarInversion(w http.ResponseWriter, r *http.Request) {

	idStr := chi.URLParam(r, "id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	var datos models.Inversion

	err = json.NewDecoder(r.Body).Decode(&datos)
	if err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}

	inversionActualizada, encontrado :=
		h.store.ActualizarInversion(id, datos)

	if !encontrado {
		http.Error(w, "Inversión no encontrada", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(inversionActualizada)
}

// BorrarInversion maneja la solicitud DELETE para eliminar una inversión por su ID.
func (h *InversionHandler) BorrarInversion(w http.ResponseWriter, r *http.Request) {

	idStr := chi.URLParam(r, "id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	borrado := h.store.BorrarInversion(id)

	if !borrado {
		http.Error(w, "Inversión no encontrada", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(map[string]string{
		"mensaje": "Inversión eliminada correctamente",
	})
}
