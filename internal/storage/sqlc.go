package storage

import (
	"context"
	"database/sql"

	"Proyecto_AWEBII/internal/models"
	"Proyecto_AWEBII/internal/storage/sqlcdb"
)

// AlmacenSQLC implementa la interfaz Almacen usando SQLC.
type AlmacenSQLC struct {
	q *sqlcdb.Queries
}

// Constructor
func NuevoAlmacenSQLC(db *sql.DB) *AlmacenSQLC {
	return &AlmacenSQLC{
		q: sqlcdb.New(db),
	}
}

// =========================
// HELPERS (IMPORTANTE)
// =========================

func ns(s sql.NullString) string {
	if s.Valid {
		return s.String
	}
	return ""
}

func ni(i sql.NullInt64) int {
	if i.Valid {
		return int(i.Int64)
	}
	return 0
}

// =========================
// MAPEOS
// =========================

func aEventoDominio(e sqlcdb.Evento) models.Evento {
	return models.Evento{
		ID:          int(e.ID),
		Nombre:      e.Nombre,
		Descripcion: ns(e.Descripcion),
		Fecha:       ns(e.Fecha),
		Lugar:       ns(e.Lugar),
		Capacidad:   ni(e.Capacidad),
		CategoriaID: ni(e.CategoriaID),
		Organizador: ns(e.Organizador),
		Estado:      ns(e.Estado),
	}
}

// =========================
// EVENTOS
// =========================

func (a *AlmacenSQLC) ListarEventos() []models.Evento {
	filas, err := a.q.ListarEventos(context.Background())
	if err != nil {
		return nil
	}

	out := make([]models.Evento, 0, len(filas))
	for _, f := range filas {
		out = append(out, aEventoDominio(f))
	}
	return out
}

func (a *AlmacenSQLC) BuscarEventoPorID(id int) (models.Evento, bool) {
	fila, err := a.q.BuscarEventoPorID(context.Background(), int64(id))
	if err != nil {
		return models.Evento{}, false
	}

	return aEventoDominio(fila), true
}

func (a *AlmacenSQLC) CrearEvento(e models.Evento) models.Evento {
	fila, err := a.q.CrearEvento(
		context.Background(),
		sqlcdb.CrearEventoParams{
			Nombre:      e.Nombre,
			Descripcion: sql.NullString{String: e.Descripcion, Valid: e.Descripcion != ""},
			Fecha:       sql.NullString{String: e.Fecha, Valid: e.Fecha != ""},
			Lugar:       sql.NullString{String: e.Lugar, Valid: e.Lugar != ""},
			Capacidad:   sql.NullInt64{Int64: int64(e.Capacidad), Valid: true},
			CategoriaID: sql.NullInt64{Int64: int64(e.CategoriaID), Valid: true},
			Organizador: sql.NullString{String: e.Organizador, Valid: e.Organizador != ""},
			Estado:      sql.NullString{String: e.Estado, Valid: e.Estado != ""},
		},
	)

	if err != nil {
		return models.Evento{}
	}

	return aEventoDominio(fila)
}

func (a *AlmacenSQLC) ActualizarEvento(id int, datos models.Evento) (models.Evento, bool) {
	fila, err := a.q.ActualizarEvento(
		context.Background(),
		sqlcdb.ActualizarEventoParams{
			ID:          int64(id),
			Nombre:      datos.Nombre,
			Descripcion: sql.NullString{String: datos.Descripcion, Valid: datos.Descripcion != ""},
			Fecha:       sql.NullString{String: datos.Fecha, Valid: datos.Fecha != ""},
			Lugar:       sql.NullString{String: datos.Lugar, Valid: datos.Lugar != ""},
			Capacidad:   sql.NullInt64{Int64: int64(datos.Capacidad), Valid: true},
			CategoriaID: sql.NullInt64{Int64: int64(datos.CategoriaID), Valid: true},
			Organizador: sql.NullString{String: datos.Organizador, Valid: datos.Organizador != ""},
			Estado:      sql.NullString{String: datos.Estado, Valid: datos.Estado != ""},
		},
	)

	if err != nil {
		return models.Evento{}, false
	}

	return aEventoDominio(fila), true
}

func (a *AlmacenSQLC) BorrarEvento(id int) bool {
	res, err := a.q.BorrarEvento(context.Background(), int64(id))
	if err != nil {
		return false
	}
	return res > 0
}
