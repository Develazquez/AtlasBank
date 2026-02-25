package repository

import (
	"banco-api/Cuenta/domain/entities"

	"github.com/google/uuid"
)

type ICuentaRepository interface {
	Create(cuenta *entities.Cuenta) (uuid.UUID, error)
	GetByID(id uuid.UUID) (*entities.Cuenta, error)
	GetAll() ([]*entities.Cuenta, error)
	Update(cuenta *entities.Cuenta) error
	Delete(id uuid.UUID) error
	GetByNumeroCuenta(numero string) (*entities.Cuenta, error)
	GetCuentasByUsuario(idUsuario uuid.UUID) ([]*entities.Cuenta, error)
}
