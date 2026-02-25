package application

import (
	"banco-api/Banco/domain/entities"
	"banco-api/Banco/domain/repository"

	"github.com/google/uuid"
)

type GetBancoUseCase struct {
	repo repository.IBancoRepository
}

func NewGetBancoUseCase(repo repository.IBancoRepository) *GetBancoUseCase {
	return &GetBancoUseCase{repo: repo}
}

func (uc *GetBancoUseCase) Execute(id uuid.UUID) (*entities.Banco, error) {
	banco, err := uc.repo.GetByID(id)
	if err != nil {
		return nil, err
	}

	if banco == nil {
		return nil, entities.ErrBancoNoEncontrado
	}

	return banco, nil
}
