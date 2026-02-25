package application

import (
	"banco-api/Usuario/domain/entities"
	"banco-api/Usuario/domain/repository"

	"github.com/google/uuid"
)

type DeleteUsuarioUseCase struct {
	repo repository.IUsuarioRepository
}

func NewDeleteUsuarioUseCase(repo repository.IUsuarioRepository) *DeleteUsuarioUseCase {
	return &DeleteUsuarioUseCase{repo: repo}
}

func (uc *DeleteUsuarioUseCase) Execute(id uuid.UUID) error {
	existente, err := uc.repo.GetByID(id)
	if err != nil || existente == nil {
		return entities.ErrUsuarioNoEncontrado
	}

	return uc.repo.Delete(id)
}
