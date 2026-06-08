package handlers

//endpoints
import (
	"encoding/json"
	"net/http"

	"Proyecto_AWEBII/internal/models"
	"Proyecto_AWEBII/internal/services"
	"fmt"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func ObtenerEventos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	eventos := services.ObtenerEventos()

	json.NewEncoder(w).Encode(eventos)
}

// ObtenerEventos()
// ObtenerEvento()
// ActualizarEvento()
// EliminarEvento()
func CrearEvento(w http.ResponseWriter, r *http.Request) {

	var evento models.Evento

	err := json.NewDecoder(r.Body).Decode(&evento)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Printf("%+v\n", evento)

	services.CrearEvento(evento)

	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(evento)
}
func ObtenerEvento(w http.ResponseWriter, r *http.Request) {

	id, _ := strconv.Atoi(chi.URLParam(r, "id"))

	evento, err := services.ObtenerEventoPorID(id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(evento)
}

func ActualizarEvento(w http.ResponseWriter, r *http.Request) {

	id, _ := strconv.Atoi(chi.URLParam(r, "id"))

	var evento models.Evento

	err := json.NewDecoder(r.Body).Decode(&evento)

	if err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}

	err = services.ActualizarEvento(id, evento)

	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(evento)
}

func EliminarEvento(w http.ResponseWriter, r *http.Request) {

	id, _ := strconv.Atoi(chi.URLParam(r, "id"))

	err := services.EliminarEvento(id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Write([]byte("Evento eliminado"))
}
