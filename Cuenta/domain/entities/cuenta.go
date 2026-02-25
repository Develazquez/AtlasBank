package entities

import (
	"time"

	"github.com/google/uuid"
)

// Enums for Cuenta
type TipoCuenta string
type EstadoCuenta string
type Moneda string
type TipoTarjeta string

const (
	// TipoCuenta
	AHORRO    TipoCuenta = "AHORRO"
	CORRIENTE TipoCuenta = "CORRIENTE"
	NOMINA    TipoCuenta = "NOMINA"
	INVERSION TipoCuenta = "INVERSION"

	// EstadoCuenta
	ACTIVA    EstadoCuenta = "ACTIVA"
	INACTIVA  EstadoCuenta = "INACTIVA"
	BLOQUEADA EstadoCuenta = "BLOQUEADA"
	CERRADA   EstadoCuenta = "CERRADA"

	// Moneda
	MXN Moneda = "MXN"
	USD Moneda = "USD"
	EUR Moneda = "EUR"

	// TipoTarjeta
	CLASSIC  TipoTarjeta = "CLASSIC"
	GOLD     TipoTarjeta = "GOLD"
	BLACK    TipoTarjeta = "BLACK"
	PLATINUM TipoTarjeta = "PLATINUM"
)

type Cuenta struct {
	ID                    uuid.UUID    `json:"id" gorm:"type:uuid;primaryKey"`
	UsuarioID             uuid.UUID    `json:"usuario_id" gorm:"type:uuid;not null" binding:"required"`
	BancoID               uuid.UUID    `json:"banco_id" gorm:"type:uuid;not null" binding:"required"`
	NumeroCuenta          string       `json:"numero_cuenta" binding:"required,min=5" gorm:"type:varchar(20);unique;not null"`
	NumeroTarjeta         string       `json:"numero_tarjeta" gorm:"type:varchar(16);unique"`
	UltimosDigitosTarjeta string       `json:"ultimos_digitos_tarjeta" gorm:"type:varchar(4);generatedColumn:STORED;->"`
	CLABE                 string       `json:"clabe" gorm:"type:varchar(18);unique"`
	IBAN                  string       `json:"iban" gorm:"type:varchar(34);unique"`
	Tipo                  TipoCuenta   `json:"tipo" gorm:"type:tipo_cuenta_enum;default:'AHORRO';not null"`
	TipoTarjeta           TipoTarjeta  `json:"tipo_tarjeta" gorm:"type:tipo_tarjeta_enum;default:'CLASSIC';not null"`
	NombreTarjeta         string       `json:"nombre_tarjeta" gorm:"type:varchar(50)"`
	FechaExpiracion       string       `json:"fecha_expiracion" gorm:"type:varchar(5)"`
	Moneda                Moneda       `json:"moneda" gorm:"type:moneda_enum;default:'MXN';not null"`
	Saldo                 float64      `json:"saldo" gorm:"type:numeric(18,2);default:0.00;not null"`
	SaldoBloqueado        float64      `json:"saldo_bloqueado" gorm:"type:numeric(18,2);default:0.00;not null"`
	LimiteRetiroDiario    float64      `json:"limite_retiro_diario" gorm:"type:numeric(18,2);default:10000.00;not null"`
	Estado                EstadoCuenta `json:"estado" gorm:"type:estado_cuenta_enum;default:'ACTIVA';not null"`
	FechaApertura         time.Time    `json:"fecha_apertura" gorm:"type:date;default:CURRENT_DATE;not null"`
	FechaCierre           *time.Time   `json:"fecha_cierre" gorm:"type:date"`
	CreatedAt             time.Time    `json:"created_at" gorm:"autoCreateTime:milli"`
	UpdatedAt             time.Time    `json:"updated_at" gorm:"autoUpdateTime:milli"`
}

func (c *Cuenta) Validar() error {
	if c.NumeroCuenta == "" {
		return ErrorCuenta{Mensaje: "El número de cuenta es requerido"}
	}
	if c.Tipo != AHORRO && c.Tipo != CORRIENTE && c.Tipo != NOMINA && c.Tipo != INVERSION {
		return ErrorCuenta{Mensaje: "El tipo de cuenta no es válido"}
	}
	if c.Saldo < 0 {
		return ErrorCuenta{Mensaje: "El saldo no puede ser negativo"}
	}
	if c.UsuarioID == uuid.Nil {
		return ErrorCuenta{Mensaje: "El ID del usuario es requerido"}
	}
	if c.BancoID == uuid.Nil {
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
	ErrCuentaBloqueada      = ErrorCuenta{Mensaje: "La cuenta está bloqueada"}
	ErrCuentaCerrada        = ErrorCuenta{Mensaje: "La cuenta está cerrada"}
)
