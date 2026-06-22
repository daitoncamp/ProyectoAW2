--------------------
-- MODULO EVENTOS
--------------------

CREATE TABLE categorias_evento (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    nombre TEXT NOT NULL,
    descripcion TEXT
);

CREATE TABLE eventos (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    nombre TEXT NOT NULL,
    descripcion TEXT NOT NULL,
    fecha TEXT NOT NULL,
    lugar TEXT,
    capacidad INTEGER,
    categoria_id INTEGER,
    organizador TEXT,
    estado TEXT,

    FOREIGN KEY (categoria_id) REFERENCES categorias_evento(id)
);

CREATE TABLE asistencias (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    evento_id INTEGER NOT NULL,
    estudiante_id INTEGER NOT NULL,
    presente INTEGER DEFAULT 0,
    hora_registro TEXT,

    FOREIGN KEY (evento_id) REFERENCES eventos(id)
);