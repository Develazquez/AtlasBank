package entities

import (
	"database/sql/driver"
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

// Enums for Transaccion
type TipoTransaccion string
type EstadoTransaccion string

const (
	// TipoTransaccion
	DEPOSITO              TipoTransaccion = "DEPOSITO"
	RETIRO                TipoTransaccion = "RETIRO"
	TRANSFERENCIA_INTERNA TipoTransaccion = "TRANSFERENCIA_INTERNA"
	TRANSFERENCIA_EXTERNA TipoTransaccion = "TRANSFERENCIA_EXTERNA"
	PAGO_SERVICIO         TipoTransaccion = "PAGO_SERVICIO"
	COMISION              TipoTransaccion = "COMISION"
	INTERES               TipoTransaccion = "INTERES"

	// EstadoTransaccion
	PENDIENTE  EstadoTransaccion = "PENDIENTE"
	COMPLETADA EstadoTransaccion = "COMPLETADA"
	FALLIDA    EstadoTransaccion = "FALLIDA"
	REVERTIDA  EstadoTransaccion = "REVERTIDA"
	CANCELADA  EstadoTransaccion = "CANCELADA"
)

type Canal string

const (
	APP      Canal = "APP"
	WEB      Canal = "WEB"
	ATM      Canal = "ATM"
	SUCURSAL Canal = "SUCURSAL"
	API      Canal = "API"
)

type Metadata map[string]interface{}

func (m Metadata) Scan(value interface{}) error {
	bytes := value.([]byte)
	return json.Unmarshal(bytes, &m)
}

func (m Metadata) Value() (driver.Value, error) {
	return json.Marshal(m)
}

type Transaccion struct {
	ID              uuid.UUID         `json:"id" gorm:"type:uuid;primaryKey"`
	CuentaOrigenID  *uuid.UUID        `json:"cuenta_origen_id" gorm:"type:uuid"`
	CuentaDestinoID *uuid.UUID        `json:"cuenta_destino_id" gorm:"type:uuid"`
	Tipo            TipoTransaccion   `json:"tipo" gorm:"type:tipo_transaccion_enum;not null"`
	Estado          EstadoTransaccion `json:"estado" gorm:"type:estado_transaccion_enum;default:'PENDIENTE';not null"`
	Monto           float64           `json:"monto" binding:"required,gt=0" gorm:"type:numeric(18,2);not null"`
	Moneda          string            `json:"moneda" gorm:"type:moneda_enum;default:'MXN';not null"`
	Comision        float64           `json:"comision" gorm:"type:numeric(18,2);default:0.00;not null"`
	Concepto        string            `json:"concepto" gorm:"type:varchar(200)"`
	Descripcion     string            `json:"descripcion" gorm:"type:text"`
	Referencia      string            `json:"referencia" gorm:"type:varchar(100);unique;not null;default:uuid_generate_v4()"`
	IPOrigen        string            `json:"ip_origen" gorm:"type:inet"`
	Canal           Canal             `json:"canal" gorm:"type:varchar(30);default:'APP';not null"`
	Metadata        Metadata          `json:"metadata" gorm:"type:jsonb"`
	ProcesadoAt     *time.Time        `json:"procesado_at" gorm:"type:timestamptz"`
	CreatedAt       time.Time         `json:"created_at" gorm:"autoCreateTime:milli"`
	UpdatedAt       time.Time         `json:"updated_at" gorm:"autoUpdateTime:milli"`
}

// TableName especifica el nombre de la tabla en PostgreSQL
func (Transaccion) TableName() string {
	return "transacciones"
}

func (t *Transaccion) Validar() error {
	tiposValidos := map[TipoTransaccion]bool{
		DEPOSITO:              true,
		RETIRO:                true,
		TRANSFERENCIA_INTERNA: true,
		TRANSFERENCIA_EXTERNA: true,
		PAGO_SERVICIO:         true,
		COMISION:              true,
		INTERES:               true,
	}

	if !tiposValidos[t.Tipo] {
		return ErrorTransaccion{Mensaje: "Tipo de transacción inválido"}
	}

	if t.Monto <= 0 {
		return ErrorTransaccion{Mensaje: "El monto debe ser mayor a 0"}
	}

	if t.Tipo == DEPOSITO && t.CuentaDestinoID == nil {
		return ErrorTransaccion{Mensaje: "La cuenta destino es requerida para depósitos"}
	}

	if t.Tipo == RETIRO && t.CuentaOrigenID == nil {
		return ErrorTransaccion{Mensaje: "La cuenta origen es requerida para retiros"}
	}

	if (t.Tipo == TRANSFERENCIA_INTERNA || t.Tipo == TRANSFERENCIA_EXTERNA) &&
		(t.CuentaOrigenID == nil || t.CuentaDestinoID == nil) {
		return ErrorTransaccion{Mensaje: "Las cuentas origen y destino son requeridas para transferencias"}
	}

	return nil
}

type ErrorTransaccion struct {
	Mensaje string
}

func (e ErrorTransaccion) Error() string {
	return e.Mensaje
}

var (
	ErrTransaccionNoEncontrada = ErrorTransaccion{Mensaje: "La transacción no fue encontrada"}
	ErrSaldoInsuficiente       = ErrorTransaccion{Mensaje: "Saldo insuficiente para la transacción"}
	ErrCuentaOrigen            = ErrorTransaccion{Mensaje: "La cuenta origen no existe"}
	ErrCuentaDestino           = ErrorTransaccion{Mensaje: "La cuenta destino no existe"}
	ErrTransaccionFallida      = ErrorTransaccion{Mensaje: "La transacción falló"}
)
