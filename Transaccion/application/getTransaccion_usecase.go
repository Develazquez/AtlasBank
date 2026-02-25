package application

import (
	"banco-api/Transaccion/domain/entities"
	"banco-api/Transaccion/domain/repository"

	"github.com/google/uuid"
)

type GetTransaccionUseCase struct {
	repo repository.ITransaccionRepository
}

func NewGetTransaccionUseCase(repo repository.ITransaccionRepository) *GetTransaccionUseCase {
	return &GetTransaccionUseCase{repo: repo}
}

func (uc *GetTransaccionUseCase) Execute(id uuid.UUID) (*entities.Transaccion, error) {
	transaccion, err := uc.repo.GetByID(id)
	if err != nil {
		return nil, err
	}
	if transaccion == nil {
		return nil, entities.ErrTransaccionNoEncontrada
	}
	return transaccion, nil
}
