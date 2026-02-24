package entities

import "time"

type Usuario struct {
	IDUsuario       int        `json:"id_usuario"`
	Nombre          string     `json:"nombre" binding:"required,min=2,max=100"`
	Apellido        string     `json:"apellido" binding:"required,min=2,max=100"`
	Email           string     `json:"email" binding:"required,email"`
	Telefono        string     `json:"telefono" binding:"max=20"`
	FechaNacimiento *time.Time `json:"fecha_nacimiento"`
	Password        string     `json:"password"` 
	CreatedAt       time.Time  `json:"created_at"`
	UpdatedAt       time.Time  `json:"updated_at"`
}

func (u *Usuario) Validar() error {
	if u.Nombre == "" {
		return ErrorUsuario{Mensaje: "El nombre es requerido"}
	}
	if u.Apellido == "" {
		return ErrorUsuario{Mensaje: "El apellido es requerido"}
	}
	if u.Email == "" {
		return ErrorUsuario{Mensaje: "El email es requerido"}
	}
	return nil
}

type ErrorUsuario struct {
	Mensaje string
}

func (e ErrorUsuario) Error() string {
	return e.Mensaje
}

var (
	ErrUsuarioNoEncontrado = ErrorUsuario{Mensaje: "El usuario no fue encontrado"}
	ErrEmailYaExiste       = ErrorUsuario{Mensaje: "El email ya existe"}
)
