package repository

import (
	"banco-api/Banco/domain/entities"

	"github.com/google/uuid"
)

type IBancoRepository interface {
	Create(banco *entities.Banco) (*entities.Banco, error)
	GetByID(id uuid.UUID) (*entities.Banco, error)
	GetAll() ([]*entities.Banco, error)
	Update(banco *entities.Banco) error
	Delete(id uuid.UUID) error
	GetByName(nombre string) (*entities.Banco, error)
	GetByCodigoSwift(codigo string) (*entities.Banco, error)
}
