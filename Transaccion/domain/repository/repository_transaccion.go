package repository

import (
	"banco-api/Transaccion/domain/entities"

	"github.com/google/uuid"
)

type ITransaccionRepository interface {
	Create(transaccion *entities.Transaccion) (uuid.UUID, error)
	GetByID(id uuid.UUID) (*entities.Transaccion, error)
	GetAll() ([]*entities.Transaccion, error)
	Delete(id uuid.UUID) error
	GetTransactionsByCuenta(idCuenta uuid.UUID) ([]*entities.Transaccion, error)
}
