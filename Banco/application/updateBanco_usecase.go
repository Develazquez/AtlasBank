package application

import (
	"banco-api/Banco/domain/entities"
	"banco-api/Banco/domain/repository"
)

type UpdateBancoUseCase struct {
	repo repository.IBancoRepository
}

func NewUpdateBancoUseCase(repo repository.IBancoRepository) *UpdateBancoUseCase {
	return &UpdateBancoUseCase{repo: repo}
}

func (uc *UpdateBancoUseCase) Execute(banco *entities.Banco) error {
	if err := banco.Validar(); err != nil {
		return err
	}

	existente, err := uc.repo.GetByID(banco.ID)
	if err != nil || existente == nil {
		return entities.ErrBancoNoEncontrado
	}

	return uc.repo.Update(banco)
}
