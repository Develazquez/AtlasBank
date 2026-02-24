package application

import (
	"banco-api/Usuario/domain/entities"
	"banco-api/Usuario/domain/repository"
)

type DeleteUsuarioUseCase struct {
	repo repository.IUsuarioRepository
}

func NewDeleteUsuarioUseCase(repo repository.IUsuarioRepository) *DeleteUsuarioUseCase {
	return &DeleteUsuarioUseCase{repo: repo}
}

func (uc *DeleteUsuarioUseCase) Execute(id int) error {
	existente, err := uc.repo.GetByID(id)
	if err != nil || existente == nil {
		return entities.ErrUsuarioNoEncontrado
	}

	return uc.repo.Delete(id)
}
