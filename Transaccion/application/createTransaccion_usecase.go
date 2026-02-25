package application

import (
	"banco-api/Transaccion/domain/entities"
	"banco-api/Transaccion/domain/repository"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CreateTransaccionUseCase struct {
	repo repository.ITransaccionRepository
	db   *gorm.DB
}

func NewCreateTransaccionUseCase(repo repository.ITransaccionRepository, db *gorm.DB) *CreateTransaccionUseCase {
	return &CreateTransaccionUseCase{
		repo: repo,
		db:   db,
	}
}

func (uc *CreateTransaccionUseCase) Execute(transaccion *entities.Transaccion) (uuid.UUID, error) {
	if err := transaccion.Validar(); err != nil {
		return uuid.Nil, err
	}

	return uc.repo.Create(transaccion)
}
