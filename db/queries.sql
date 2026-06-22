
-- EVENTOS

-- name: ListarEventos :many
SELECT * FROM eventos;

-- name: BuscarEventoPorID :one
SELECT * FROM eventos WHERE id = ?;

-- name: CrearEvento :one
INSERT INTO eventos (
  nombre, descripcion, fecha, lugar, capacidad, categoria_id, organizador, estado
) VALUES (?, ?, ?, ?, ?, ?, ?, ?)
RETURNING *;

-- name: ActualizarEvento :one
UPDATE eventos
SET
  nombre = ?,
  descripcion = ?,
  fecha = ?,
  lugar = ?,
  capacidad = ?,
  categoria_id = ?,
  organizador = ?,
  estado = ?
WHERE id = ?
RETURNING *;

-- name: BorrarEvento :execrows
DELETE FROM eventos WHERE id = ?;