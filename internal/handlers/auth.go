package handlers

import (
	"encoding/json"
	"net/http"

	"Proyecto_AWEBII/internal/services"
)

type AuthHandler struct {
	auth *services.AuthService
}

func NewAuthHandler(auth *services.AuthService) *AuthHandler {
	return &AuthHandler{
		auth: auth,
	}
}

// Credenciales representa los datos enviados para registrarse o iniciar sesión.
type Credenciales struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Registrar crea un nuevo usuario.
func (h *AuthHandler) Registrar(w http.ResponseWriter, r *http.Request) {

	var creds Credenciales

	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}

	usuario, err := h.auth.Register(creds.Email, creds.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(usuario)
}

// Login autentica un usuario.
func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {

	var creds Credenciales

	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}

	token, err := h.auth.Login(creds.Email, creds.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(map[string]string{
		"token": token,
	})
}
