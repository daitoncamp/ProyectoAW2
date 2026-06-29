
-- EVENTOS

-- name: ListarEventos :many
SELECT * FROM eventos ORDER BY id;

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


--------------------
--INVERSION--
----------------------

-- name: ListarInversiones :many
SELECT *
FROM inversiones
ORDER BY id;

-- name: BuscarInversionPorID :one
SELECT *
FROM inversiones
WHERE id = ?;


-- name: CrearInversion :one
INSERT INTO inversiones (
    nombre,
    monto_inicial,
    monto_actual,
    rendimiento_esperado,
    estado,
    tipo_inversion_id,
    destino_inversion_id
)
VALUES (
    ?, ?, ?, ?, ?, ?, ?
)
RETURNING *;


-- name: ActualizarInversion :exec
UPDATE inversiones
SET
    nombre = ?,
    monto_inicial = ?,
    monto_actual = ?,
    rendimiento_esperado = ?,
    estado = ?,
    tipo_inversion_id = ?,
    destino_inversion_id = ?
WHERE id = ?;



-- name: BorrarInversion :exec
DELETE FROM inversiones
WHERE id = ?;



-- name: ListarAportesPorInversion :many
SELECT *
FROM aportes
WHERE inversion_id = ?
ORDER BY id;



-- name: CrearAporte :one
INSERT INTO aportes (
    inversion_id,
    nombre,
    monto
)
VALUES (
    ?, ?, ?
)
RETURNING *;


-- name: ListarTiposInversion :many
SELECT *
FROM tipos_inversion
ORDER BY id;


-- name: ListarDestinosInversion :many
SELECT *
FROM destinos_inversion
ORDER BY id;
