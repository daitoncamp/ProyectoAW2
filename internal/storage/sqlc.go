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

// Verificación en tiempo de compilación.
//var _ Almacen = (*AlmacenSQLC)(nil)
