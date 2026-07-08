package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"Proyecto_AWEBII/internal/models"
	"Proyecto_AWEBII/internal/services"

	"github.com/go-chi/chi/v5"
)

// LISTAR ESTUDIANTES
func ObtenerEstudiantes(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	estudiantes := services.ObtenerEstudiantes()

	if err := json.NewEncoder(w).Encode(estudiantes); err != nil {
		http.Error(w, "Error al generar la respuesta", http.StatusInternalServerError)
		return
	}
}

// CREAR ESTUDIANTE
func CrearEstudiante(w http.ResponseWriter, r *http.Request) {

	var estudiante models.Estudiante

	if err := json.NewDecoder(r.Body).Decode(&estudiante); err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}

	services.CrearEstudiante(estudiante)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	if err := json.NewEncoder(w).Encode(estudiante); err != nil {
		http.Error(w, "Error al generar la respuesta", http.StatusInternalServerError)
		return
	}
}

// OBTENER ESTUDIANTE POR ID
func ObtenerEstudiante(w http.ResponseWriter, r *http.Request) {

	id, err := strconv.Atoi(chi.URLParam(r, "id"))

	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	estudiante, err := services.ObtenerEstudiantePorID(id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(estudiante); err != nil {
		http.Error(w, "Error al generar la respuesta", http.StatusInternalServerError)
		return
	}
}

// ACTUALIZAR ESTUDIANTE
func ActualizarEstudiante(w http.ResponseWriter, r *http.Request) {

	id, err := strconv.Atoi(chi.URLParam(r, "id"))

	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	var estudiante models.Estudiante

	if err := json.NewDecoder(r.Body).Decode(&estudiante); err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}

	err = services.ActualizarEstudiante(id, estudiante)

	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(estudiante); err != nil {
		http.Error(w, "Error al generar la respuesta", http.StatusInternalServerError)
		return
	}
}

// ELIMINAR ESTUDIANTE
func EliminarEstudiante(w http.ResponseWriter, r *http.Request) {

	id, err := strconv.Atoi(chi.URLParam(r, "id"))

	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	err = services.EliminarEstudiante(id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}