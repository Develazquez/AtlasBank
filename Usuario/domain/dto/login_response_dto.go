package dto

import (
	"database/sql/driver"
	"encoding/json"
)

// CardDTO representa la información de la tarjeta
type CardDTO struct {
	NameCard   string `json:"name_card"`   // Tipo de tarjeta (CLASSIC, GOLD, BLACK, PLATINUM)
	NumCard    string `json:"num_card"`    // Últimos 4 dígitos
	Expires    string `json:"expires"`     // Fecha de expiración (MM/YY)
	CardHolder string `json:"card_holder"` // Nombre del titular
}

// RecentlyInfDTO representa la información de la última transacción
type RecentlyInfDTO struct {
	Concept     string  `json:"concept"`      // Concepto de la transacción
	TypeInf     string  `json:"type_inf"`     // '+' para entrada, '-' para salida
	DayTransfer string  `json:"day_transfer"` // Fecha y hora de la transacción
	Mount       float64 `json:"mount"`        // Monto de la transacción
}

// UserDashboardDTO representa la respuesta del login con datos del dashboard
// Estructura que coincide con v_usuario_dashboard
type UserDashboardDTO struct {
	ID          string          `json:"id"`           // ID del usuario como string
	Name        string          `json:"name"`         // Nombre completo del usuario
	Wallet      float64         `json:"wallet"`       // Saldo de la cuenta
	Card        CardDTO         `json:"card"`         // Información de la tarjeta
	RecentlyInf *RecentlyInfDTO `json:"recently_inf"` // Última transacción completada
}

// Scan implementa el interfaz sql.Scanner para permitir que GORM escanee JSON
func (c *CardDTO) Scan(value interface{}) error {
	bytes, _ := value.([]byte)
	return json.Unmarshal(bytes, &c)
}

// Value implementa el interfaz driver.Valuer para permitir que GORM escriba JSON
func (c CardDTO) Value() (driver.Value, error) {
	return json.Marshal(c)
}

// Scan implementa el interfaz sql.Scanner para permitir que GORM escanee JSON
func (r *RecentlyInfDTO) Scan(value interface{}) error {
	if value == nil {
		return nil
	}
	bytes, _ := value.([]byte)
	return json.Unmarshal(bytes, &r)
}

// Value implementa el interfaz driver.Valuer para permitir que GORM escriba JSON
func (r RecentlyInfDTO) Value() (driver.Value, error) {
	return json.Marshal(r)
}
