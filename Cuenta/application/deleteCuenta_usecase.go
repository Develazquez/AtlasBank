package application

import (
	"banco-api/Cuenta/domain/entities"
	"banco-api/Cuenta/domain/repository"

	"github.com/google/uuid"
)

type DeleteCuentaUseCase struct {
	repo repository.ICuentaRepository
}

func NewDeleteCuentaUseCase(repo repository.ICuentaRepository) *DeleteCuentaUseCase {
	return &DeleteCuentaUseCase{repo: repo}
}

func (uc *DeleteCuentaUseCase) Execute(id uuid.UUID) error {
	existente, err := uc.repo.GetByID(id)
	if err != nil || existente == nil {
		return entities.ErrCuentaNoEncontrada
	}

	return uc.repo.Delete(id)
}
