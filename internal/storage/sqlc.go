package storage

import (
	"context"
	"database/sql"

	"Proyecto_AWEBII/internal/models"
	"Proyecto_AWEBII/internal/storage/sqlcdb"
)

// AlmacenSQLC implementa la interfaz Almacen usando sqlc.
type AlmacenSQLC struct {
	q *sqlcdb.Queries
}

// NuevoAlmacenSQLC envuelve una conexión *sql.DB ya abierta.
func NuevoAlmacenSQLC(db *sql.DB) *AlmacenSQLC {
	return &AlmacenSQLC{
		q: sqlcdb.New(db),
	}
}

// =========================================================
// MAPEO sqlc <-> dominio
// =========================================================

func aInversionDominio(i sqlcdb.Inversion) models.Inversion {
	return models.Inversion{
		ID:                  int(i.ID),
		Nombre:              i.Nombre,
		MontoInicial:        i.MontoInicial,
		MontoActual:         i.MontoActual,
		RendimientoEsperado: i.RendimientoEsperado,
		Estado:              i.Estado,
		TipoInversionID:     int(i.TipoInversionID),
		DestinoInversionID:  int(i.DestinoInversionID),
	}
}

// =========================================================
// INVERSIONES
// =========================================================

func (a *AlmacenSQLC) ListarInversiones() []models.Inversion {

	filas, err := a.q.ListarInversiones(context.Background())
	if err != nil {
		return nil
	}

	out := make([]models.Inversion, 0, len(filas))

	for _, f := range filas {
		out = append(out, aInversionDominio(f))
	}

	return out
}

func (a *AlmacenSQLC) BuscarInversionPorID(id int) (models.Inversion, bool) {

	fila, err := a.q.BuscarInversionPorID(
		context.Background(),
		int64(id),
	)

	if err != nil {
		return models.Inversion{}, false
	}

	return aInversionDominio(fila), true
}

func (a *AlmacenSQLC) CrearInversion(
	i models.Inversion,
) models.Inversion {

	fila, err := a.q.CrearInversion(
		context.Background(),
		sqlcdb.CrearInversionParams{
			Nombre:              i.Nombre,
			MontoInicial:        i.MontoInicial,
			MontoActual:         i.MontoActual,
			RendimientoEsperado: i.RendimientoEsperado,
			Estado:              i.Estado,
			TipoInversionID:     int64(i.TipoInversionID),
			DestinoInversionID:  int64(i.DestinoInversionID),
		},
	)

	if err != nil {
		return models.Inversion{}
	}

	return aInversionDominio(fila)
}

func (a *AlmacenSQLC) ActualizarInversion(
	id int,
	datos models.Inversion,
) (models.Inversion, bool) {

	fila, err := a.q.ActualizarInversion(
		context.Background(),
		sqlcdb.ActualizarInversionParams{
			Nombre:              datos.Nombre,
			MontoInicial:        datos.MontoInicial,
			MontoActual:         datos.MontoActual,
			RendimientoEsperado: datos.RendimientoEsperado,
			Estado:              datos.Estado,
			TipoInversionID:     int64(datos.TipoInversionID),
			DestinoInversionID:  int64(datos.DestinoInversionID),
			ID:                  int64(id),
		},
	)

	if err != nil {
		return models.Inversion{}, false
	}

	return aInversionDominio(fila), true
}

func (a *AlmacenSQLC) BorrarInversion(id int) bool {

	filas, err := a.q.BorrarInversion(
		context.Background(),
		int64(id),
	)

	if err != nil {
		return false
	}

	return filas > 0
}

func aEventoDominio(e sqlcdb.Evento) models.Evento {
	return models.Evento{
		ID:          int(e.ID),
		Nombre:      e.Nombre,
		Descripcion: e.Descripcion.String,
		Fecha:       e.Fecha.String,
		Lugar:       e.Lugar.String,
		Capacidad:   int(e.Capacidad.Int64),
		CategoriaID: int(e.CategoriaID.Int64),
		Organizador: e.Organizador.String,
		Estado:      e.Estado.String,
	}
}

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
			Nombre: e.Nombre,

			Descripcion: sql.NullString{
				String: e.Descripcion,
				Valid:  true,
			},

			Fecha: sql.NullString{
				String: e.Fecha,
				Valid:  true,
			},

			Lugar: sql.NullString{
				String: e.Lugar,
				Valid:  true,
			},

			Capacidad: sql.NullInt64{
				Int64: int64(e.Capacidad),
				Valid: true,
			},

			CategoriaID: sql.NullInt64{
				Int64: int64(e.CategoriaID),
				Valid: true,
			},

			Organizador: sql.NullString{
				String: e.Organizador,
				Valid:  true,
			},

			Estado: sql.NullString{
				String: e.Estado,
				Valid:  true,
			},
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
			ID:     int64(id),
			Nombre: datos.Nombre,

			Descripcion: sql.NullString{
				String: datos.Descripcion,
				Valid:  true,
			},

			Fecha: sql.NullString{
				String: datos.Fecha,
				Valid:  true,
			},

			Lugar: sql.NullString{
				String: datos.Lugar,
				Valid:  true,
			},

			Capacidad: sql.NullInt64{
				Int64: int64(datos.Capacidad),
				Valid: true,
			},

			CategoriaID: sql.NullInt64{
				Int64: int64(datos.CategoriaID),
				Valid: true,
			},

			Organizador: sql.NullString{
				String: datos.Organizador,
				Valid:  true,
			},

			Estado: sql.NullString{
				String: datos.Estado,
				Valid:  true,
			},
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

// Verificación en tiempo de compilación.
//var _ Almacen = (*AlmacenSQLC)(nil)
