package entities

import "time"

type Banco struct {
	IDBanco   int       `json:"id_banco"`
	Nombre    string    `json:"nombre" binding:"required,min=3,max=100"`
	Direccion string    `json:"direccion" binding:"max=200"`
	Telefono  string    `json:"telefono" binding:"max=20"`
	Email     string    `json:"email" binding:"email"`
	Activo    bool      `json:"activo"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}


func (b *Banco) Validar() error {
	if b.Nombre == "" {
		return ErrBancoNombreRequerido
	}
	if len(b.Nombre) > 100 {
		return ErrBancoNombreLargo
	}
	return nil
}


var (
	ErrBancoNombreRequerido = ErrorBanco{Mensaje: "El nombre del banco es requerido"}
	ErrBancoNombreLargo     = ErrorBanco{Mensaje: "El nombre del banco no debe exceder 100 caracteres"}
	ErrBancoNoEncontrado    = ErrorBanco{Mensaje: "El banco no fue encontrado"}
)

type ErrorBanco struct {
	Mensaje string
}

func (e ErrorBanco) Error() string {
	return e.Mensaje
}
