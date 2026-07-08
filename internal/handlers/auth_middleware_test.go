package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
)

// TestRutaProtegidaSinToken verifica que una ruta protegida
// responda 401 Unauthorized cuando no se envía el token JWT.
x

	r := chi.NewRouter()

	// Middleware que protege las rutas.
	r.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			auth := r.Header.Get("Authorization")

			if auth == "" {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}

			next.ServeHTTP(w, r)
		})
	})

	r.Get("/protegida", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	req := httptest.NewRequest(
		http.MethodGet,
		"/protegida",
		nil,
	)

	rr := httptest.NewRecorder()

	r.ServeHTTP(rr, req)

	if rr.Code != http.StatusUnauthorized {
		t.Fatalf(
			"Se esperaba %d pero llegó %d",
			http.StatusUnauthorized,
			rr.Code,
		)
	}
}
