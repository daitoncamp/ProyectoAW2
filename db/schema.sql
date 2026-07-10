
---INVERSIONES-----
CREATE TABLE tipos_inversion (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    nombre TEXT NOT NULL,
    descripcion TEXT
);

CREATE TABLE destinos_inversion (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    nombre TEXT NOT NULL,
    descripcion TEXT
);

CREATE TABLE inversiones (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    nombre TEXT NOT NULL,
    monto_inicial REAL NOT NULL,
    monto_actual REAL NOT NULL,
    rendimiento_esperado REAL NOT NULL,
    estado TEXT NOT NULL,

    tipo_inversion_id INTEGER NOT NULL,
    destino_inversion_id INTEGER NOT NULL,

    FOREIGN KEY (tipo_inversion_id)
        REFERENCES tipos_inversion(id),

    FOREIGN KEY (destino_inversion_id)
        REFERENCES destinos_inversion(id)
);

CREATE TABLE aportes (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    inversion_id INTEGER NOT NULL,
    nombre TEXT NOT NULL,
    monto REAL NOT NULL,

    FOREIGN KEY (inversion_id)
        REFERENCES inversiones(id)
);

-- MODULO EVENTOS

CREATE TABLE eventos (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    nombre TEXT NOT NULL,
    descripcion TEXT NOT NULL,
    fecha TEXT NOT NULL,
    lugar TEXT NOT NULL,
    capacidad INTEGER NOT NULL,
    categoria_id INTEGER NOT NULL,
    organizador TEXT NOT NULL,
    estado TEXT NOT NULL,

    FOREIGN KEY (categoria_id) REFERENCES categorias_evento(id)
);

CREATE TABLE categorias_evento (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    nombre TEXT NOT NULL,
    descripcion TEXT NOT NULL
);


CREATE TABLE asistencias (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    evento_id INTEGER NOT NULL,
    estudiante_id INTEGER NOT NULL,
    presente INTEGER NOT NULL DEFAULT 0,
    hora_registro TEXT NOT NULL,

    FOREIGN KEY (evento_id) REFERENCES eventos(id)
);