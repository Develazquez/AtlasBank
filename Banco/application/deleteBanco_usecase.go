package application

import (
	"banco-api/Banco/domain/entities"
	"banco-api/Banco/domain/repository"
)

type DeleteBancoUseCase struct {
	repo repository.IBancoRepository
}

func NewDeleteBancoUseCase(repo repository.IBancoRepository) *DeleteBancoUseCase {
	return &DeleteBancoUseCase{repo: repo}
}

func (uc *DeleteBancoUseCase) Execute(id int) error {
	existente, err := uc.repo.GetByID(id)
	if err != nil || existente == nil {
		return entities.ErrBancoNoEncontrado
	}

	return uc.repo.Delete(id)
}
