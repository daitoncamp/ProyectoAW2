package services

import (
	"strings"
	"time"

	"Proyecto_AWEBII/internal/models"
	"Proyecto_AWEBII/internal/storage"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

var secretoJWT = []byte("proyecto_awebii_jwt")

var duracionToken = 24 * time.Hour

// Claims contiene la información que viajará dentro del JWT.
type Claims struct {
	UsuarioID int `json:"uid"`
	jwt.RegisteredClaims
}

// AuthService contiene la lógica de autenticación.
type AuthService struct {
	repo storage.UserRepository
}

// NuevoAuthService crea una nueva instancia del servicio.
func NuevoAuthService(repo storage.UserRepository) *AuthService {
	return &AuthService{
		repo: repo,
	}
}

// Register registra un nuevo usuario.
func (s *AuthService) Register(email, password string) (models.Usuario, error) {

	email = strings.TrimSpace(strings.ToLower(email))
	password = strings.TrimSpace(password)

	if email == "" || password == "" {
		return models.Usuario{}, ErrCamposVacios
	}

	if _, existe := s.repo.BuscarUsuarioPorEmail(email); existe {
		return models.Usuario{}, ErrEmailEnUso
	}

	hash, err := bcrypt.GenerateFromPassword(
		[]byte(password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		return models.Usuario{}, err
	}

	return s.repo.CrearUsuario(models.Usuario{
		Email:        email,
		PasswordHash: string(hash),
	})
}

// Login verifica las credenciales y devuelve un JWT.
func (s *AuthService) Login(email, password string) (string, error) {

	email = strings.TrimSpace(strings.ToLower(email))

	u, existe := s.repo.BuscarUsuarioPorEmail(email)
	if !existe {
		return "", ErrCredencialesInvalidas
	}

	if err := bcrypt.CompareHashAndPassword(
		[]byte(u.PasswordHash),
		[]byte(password),
	); err != nil {
		return "", ErrCredencialesInvalidas
	}

	return s.generarToken(u)
}

// generarToken crea un JWT para el usuario.
func (s *AuthService) generarToken(u models.Usuario) (string, error) {

	claims := Claims{
		UsuarioID: u.ID,
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(duracionToken)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(secretoJWT)
}

// ValidarToken verifica un JWT y devuelve el ID del usuario.
func (s *AuthService) ValidarToken(tokenStr string) (int, error) {

	token, err := jwt.ParseWithClaims(
		tokenStr,
		&Claims{},
		func(token *jwt.Token) (interface{}, error) {

			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, ErrCredencialesInvalidas
			}

			return secretoJWT, nil
		},
	)

	if err != nil || !token.Valid {
		return 0, ErrCredencialesInvalidas
	}

	claims, ok := token.Claims.(*Claims)
	if !ok {
		return 0, ErrCredencialesInvalidas
	}

	return claims.UsuarioID, nil
}
