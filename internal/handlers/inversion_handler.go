package handlers

import (
	"encoding/json"
	"net/http"

	"Proyecto_AWEBII/internal/models"
	"Proyecto_AWEBII/internal/storage"
)

func CreateInversion(w http.ResponseWriter, r *http.Request) {

	var inversion models.Inversion

	err := json.NewDecoder(r.Body).Decode(&inversion)
	if err != nil {
		http.Error(w, "JSON invalido", http.StatusBadRequest)
		return
	}

	// Validaciones básicas
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

	// Asignar ID automático
	inversion.ID = storage.NextID
	storage.NextID++

	// Guardar en memoria
	storage.Inversiones = append(storage.Inversiones, inversion)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(inversion)
}
