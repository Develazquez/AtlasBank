package application

import (
	"banco-api/Cuenta/domain/entities"
	"banco-api/Cuenta/domain/repository"

	"github.com/google/uuid"
)

type GetCuentaUseCase struct {
	repo repository.ICuentaRepository
}

func NewGetCuentaUseCase(repo repository.ICuentaRepository) *GetCuentaUseCase {
	return &GetCuentaUseCase{repo: repo}
}

func (uc *GetCuentaUseCase) Execute(id uuid.UUID) (*entities.Cuenta, error) {
	cuenta, err := uc.repo.GetByID(id)
	if err != nil {
		return nil, err
	}
	if cuenta == nil {
		return nil, entities.ErrCuentaNoEncontrada
	}
	return cuenta, nil
}
