package middleware

import (
	"Proyecto_AWEBII/internal/services"
	"net/http"

	"context"
	"strings"
)

type contextKey string

const claveUsuarioID contextKey = "usuarioID"

func Auth(auth *services.AuthService) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			encabezado := r.Header.Get("Authorization")
			partes := strings.SplitN(encabezado, " ", 2)
			if len(partes) != 2 || partes[0] != "Bearer" {
				responderNoAutorizado(w)
				return
			}
			usuarioID, err := auth.ValidarToken(partes[1])
			if err != nil {
				responderNoAutorizado(w)
				return
			}
			ctx := context.WithValue(r.Context(), claveUsuarioID, usuarioID)
			next.ServeHTTP(w, r.WithContext(ctx))
		})

	}
}

func responderNoAutorizado(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusUnauthorized)
	_, _ = w.Write([]byte("Error: Token de autenticación no autorizado"))
}
