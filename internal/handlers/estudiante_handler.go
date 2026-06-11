package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"Proyecto_AWEBII/internal/models"
	"Proyecto_AWEBII/internal/services"

	"github.com/go-chi/chi/v5"
)

func ObtenerEstudiantes(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	estudiantes := services.ObtenerEstudiantes()

	json.NewEncoder(w).Encode(estudiantes)
}

func CrearEstudiante(w http.ResponseWriter, r *http.Request) {

	var estudiante models.Estudiante

	err := json.NewDecoder(r.Body).Decode(&estudiante)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Printf("%+v\n", estudiante)

	services.CrearEstudiante(estudiante)

	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(estudiante)
}

func ObtenerEstudiante(w http.ResponseWriter, r *http.Request) {

	id, _ := strconv.Atoi(chi.URLParam(r, "id"))

	estudiante, err := services.ObtenerEstudiantePorID(id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(estudiante)
}

func ActualizarEstudiante(w http.ResponseWriter, r *http.Request) {

	id, _ := strconv.Atoi(chi.URLParam(r, "id"))

	var estudiante models.Estudiante

	err := json.NewDecoder(r.Body).Decode(&estudiante)

	if err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}

	err = services.ActualizarEstudiante(id, estudiante)

	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(estudiante)
}

func EliminarEstudiante(w http.ResponseWriter, r *http.Request) {

	id, _ := strconv.Atoi(chi.URLParam(r, "id"))

	err := services.EliminarEstudiante(id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Write([]byte("Estudiante eliminado"))
}
