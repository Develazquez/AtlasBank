package dto

import (
	"github.com/google/uuid"
)

// CardDTO representa la información de la tarjeta
type CardDTO struct {
	NameCard   string `json:"name_card"`   // Tipo de tarjeta (CLASSIC, GOLD, BLACK, PLATINUM)
	NumCard    string `json:"num_card"`    // Últimos 4 dígitos
	Expires    string `json:"expires"`     // Fecha de expiración (MM/YY)
	CardHolder string `json:"card_holder"` // Nombre del titular
}

// UserDashboardDTO representa la respuesta del login con datos del dashboard
type UserDashboardDTO struct {
	UsuarioID      uuid.UUID `json:"usuario_id"`
	NombreCompleto string    `json:"nombre_completo"`
	Email          string    `json:"email"`
	Wallet         float64   `json:"wallet"` // Saldo de la cuenta
	Card           CardDTO   `json:"card"`   // Información de la tarjeta
	NumeroCuenta   string    `json:"numero_cuenta"`
	CLABE          string    `json:"clabe"`
	IBAN           string    `json:"iban"`
	CodigoSwift    string    `json:"codigo_swift"`
	BancoNombre    string    `json:"banco_nombre"`
	CuentaID       uuid.UUID `json:"cuenta_id"`
}
