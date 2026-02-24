package entities

import "time"

type Transaccion struct {
	IDTransaccion   int       `json:"id_transaccion"`
	TipoTransaccion string    `json:"tipo_transaccion" binding:"required,oneof=DEPOSITO RETIRO TRANSFERENCIA"`
	Monto           float64   `json:"monto" binding:"required,gt=0"`
	Fecha           time.Time `json:"fecha"`
	CuentaOrigen    *int      `json:"cuenta_origen"`
	CuentaDestino   *int      `json:"cuenta_destino"`
	Descripcion     string    `json:"descripcion" binding:"max=255"`
	CreatedAt       time.Time `json:"created_at"`
}

func (t *Transaccion) Validar() error {
	if t.TipoTransaccion != "DEPOSITO" && t.TipoTransaccion != "RETIRO" && t.TipoTransaccion != "TRANSFERENCIA" {
		return ErrorTransaccion{Mensaje: "Tipo de transacción inválido"}
	}
	if t.Monto <= 0 {
		return ErrorTransaccion{Mensaje: "El monto debe ser mayor a 0"}
	}
	if t.TipoTransaccion == "DEPOSITO" && t.CuentaDestino == nil {
		return ErrorTransaccion{Mensaje: "La cuenta destino es requerida para depósitos"}
	}
	if t.TipoTransaccion == "RETIRO" && t.CuentaOrigen == nil {
		return ErrorTransaccion{Mensaje: "La cuenta origen es requerida para retiros"}
	}
	if t.TipoTransaccion == "TRANSFERENCIA" && (t.CuentaOrigen == nil || t.CuentaDestino == nil) {
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
)
