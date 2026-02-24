package entities

import "time"

type Cuenta struct {
	IDCuenta     int       `json:"id_cuenta"`
	NumeroCuenta string    `json:"numero_cuenta" binding:"required,min=5"`
	TipoCuenta   string    `json:"tipo_cuenta" binding:"required,oneof=AHORRO CORRIENTE"`
	Saldo        float64   `json:"saldo"`
	IDUsuario    int       `json:"id_usuario" binding:"required"`
	IDBanco      int       `json:"id_banco" binding:"required"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

func (c *Cuenta) Validar() error {
	if c.NumeroCuenta == "" {
		return ErrorCuenta{Mensaje: "El número de cuenta es requerido"}
	}
	if c.TipoCuenta != "AHORRO" && c.TipoCuenta != "CORRIENTE" {
		return ErrorCuenta{Mensaje: "El tipo de cuenta debe ser AHORRO o CORRIENTE"}
	}
	if c.Saldo < 0 {
		return ErrorCuenta{Mensaje: "El saldo no puede ser negativo"}
	}
	if c.IDUsuario == 0 {
		return ErrorCuenta{Mensaje: "El ID del usuario es requerido"}
	}
	if c.IDBanco == 0 {
		return ErrorCuenta{Mensaje: "El ID del banco es requerido"}
	}
	return nil
}

type ErrorCuenta struct {
	Mensaje string
}

func (e ErrorCuenta) Error() string {
	return e.Mensaje
}

var (
	ErrCuentaNoEncontrada   = ErrorCuenta{Mensaje: "La cuenta no fue encontrada"}
	ErrNumeroCuentaYaExiste = ErrorCuenta{Mensaje: "El número de cuenta ya existe"}
	ErrSaldoInsuficiente    = ErrorCuenta{Mensaje: "Saldo insuficiente"}
)
