package config

import (
	"fmt"
	"os"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	Puerto       string
	DBDriver     string
	DBDsn        string
	RutaDB       string
	JWTSecreto   []byte
	JWTDuracion  time.Duration
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

func Cargar() Config {

	_ = godotenv.Load()

	dsn := conTexto("DB_DSN", "")

	if dsn == "" {

		host := conTexto("DB_HOST", "localhost")
		port := conTexto("DB_PORT", "5432")
		user := conTexto("DB_USER", "postgres")
		password := conTexto("DB_PASSWORD", "postgres")
		db := conTexto("DB_NAME", "asociacion")

		dsn = fmt.Sprintf(
			"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
			host,
			user,
			password,
			db,
			port,
		)
	}

	return Config{
		Puerto:       conTexto("PUERTO", ":8080"),
		DBDriver:     conTexto("DB_DRIVER", "sqlite"),
		DBDsn:        dsn,
		RutaDB:       conTexto("RUTA_DB", "proyectoaw2.db"),
		JWTSecreto:   []byte(conTexto("JWT_SECRETO", "proyectoaw2-estudiantil-dev")),
		JWTDuracion:  conDuracion("JWT_DURACION", 24*time.Hour),
		ReadTimeout:  conDuracion("HTTP_READ_TIMEOUT", 10*time.Second),
		WriteTimeout: conDuracion("HTTP_WRITE_TIMEOUT", 10*time.Second),
	}
}

func conTexto(clave, porDefecto string) string {
	if v := os.Getenv(clave); v != "" {
		return v
	}
	return porDefecto
}

func conDuracion(clave string, porDefecto time.Duration) time.Duration {
	v := os.Getenv(clave)
	if v == "" {
		return porDefecto
	}

	d, err := time.ParseDuration(v)
	if err != nil {
		return porDefecto
	}

	return d
}
