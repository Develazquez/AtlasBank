package entities

import (
	"time"

	"github.com/google/uuid"
)

type Banco struct {
	ID          uuid.UUID `json:"id" gorm:"type:uuid;primaryKey"`
	Nombre      string    `json:"nombre" binding:"required,min=3,max=100" gorm:"type:varchar(100);not null"`
	CodigoSwift string    `json:"codigo_swift" binding:"required,len=11" gorm:"type:varchar(11);unique;not null"`
	CodigoIban  string    `json:"codigo_iban" gorm:"type:varchar(34)"`
	RUC         string    `json:"ruc" binding:"required" gorm:"type:varchar(20);unique;not null"`
	Pais        string    `json:"pais" gorm:"type:varchar(60);not null;default:'México'"`
	Direccion   string    `json:"direccion" gorm:"type:text"`
	Telefono    string    `json:"telefono" gorm:"type:varchar(20)"`
	Email       string    `json:"email" binding:"email" gorm:"type:varchar(100)"`
	SitioWeb    string    `json:"sitio_web" gorm:"type:varchar(150)"`
	Activo      bool      `json:"activo" gorm:"default:true"`
	CreatedAt   time.Time `json:"created_at" gorm:"autoCreateTime:milli"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"autoUpdateTime:milli"`
}

func (b *Banco) Validar() error {
	if b.Nombre == "" {
		return ErrBancoNombreRequerido
	}
	if len(b.Nombre) > 100 {
		return ErrBancoNombreLargo
	}
	if b.CodigoSwift == "" || len(b.CodigoSwift) != 11 {
		return ErrBancoCodigoSwiftInvalido
	}
	if b.RUC == "" {
		return ErrBancoRUCRequerido
	}
	return nil
}

var (
	ErrBancoNombreRequerido     = ErrorBanco{Mensaje: "El nombre del banco es requerido"}
	ErrBancoNombreLargo         = ErrorBanco{Mensaje: "El nombre del banco no debe exceder 100 caracteres"}
	ErrBancoNoEncontrado        = ErrorBanco{Mensaje: "El banco no fue encontrado"}
	ErrBancoCodigoSwiftInvalido = ErrorBanco{Mensaje: "El código SWIFT debe tener 11 caracteres"}
	ErrBancoRUCRequerido        = ErrorBanco{Mensaje: "El RUC del banco es requerido"}
)

type ErrorBanco struct {
	Mensaje string
}

func (e ErrorBanco) Error() string {
	return e.Mensaje
}
