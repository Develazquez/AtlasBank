package entities

import (
	"time"

	"github.com/google/uuid"
)

type Usuario struct {
	ID              uuid.UUID  `json:"id" gorm:"type:uuid;primaryKey"`
	BancoID         uuid.UUID  `json:"banco_id" gorm:"type:uuid;not null" binding:"required"`
	Nombre          string     `json:"nombre" binding:"required,min=2,max=80" gorm:"type:varchar(80);not null"`
	ApellidoPaterno string     `json:"apellido_paterno" binding:"required,min=2,max=80" gorm:"type:varchar(80);not null"`
	ApellidoMaterno string     `json:"apellido_materno" gorm:"type:varchar(80)"`
	Email           string     `json:"email" binding:"required,email" gorm:"type:varchar(150);unique;not null"`
	Telefono        string     `json:"telefono" gorm:"type:varchar(20)"`
	FechaNacimiento time.Time  `json:"fecha_nacimiento" binding:"required" gorm:"type:date;not null"`
	TipoDocumento   string     `json:"tipo_documento" gorm:"type:varchar(20);default:'INE'" binding:"oneof=INE PASAPORTE CURP RFC"`
	NumeroDocumento string     `json:"numero_documento" binding:"required" gorm:"type:varchar(50);unique;not null"`
	PasswordHash    string     `json:"-" gorm:"type:text;not null"`
	Rol             string     `json:"rol" gorm:"type:varchar(20);default:'CLIENTE'" binding:"oneof=ADMIN CAJERO CLIENTE"`
	Activo          bool       `json:"activo" gorm:"default:true"`
	UltimoLogin     *time.Time `json:"ultimo_login" gorm:"type:timestamptz"`
	MemberSince     time.Time  `json:"member_since" gorm:"type:date;default:CURRENT_DATE"`
	CreatedAt       time.Time  `json:"created_at" gorm:"autoCreateTime:milli"`
	UpdatedAt       time.Time  `json:"updated_at" gorm:"autoUpdateTime:milli"`
}

func (u *Usuario) Validar() error {
	if u.Nombre == "" {
		return ErrorUsuario{Mensaje: "El nombre es requerido"}
	}
	if u.ApellidoPaterno == "" {
		return ErrorUsuario{Mensaje: "El apellido paterno es requerido"}
	}
	if u.Email == "" {
		return ErrorUsuario{Mensaje: "El email es requerido"}
	}
	if u.NumeroDocumento == "" {
		return ErrorUsuario{Mensaje: "El número de documento es requerido"}
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
	ErrDocumentoYaExiste   = ErrorUsuario{Mensaje: "El número de documento ya existe"}
)
