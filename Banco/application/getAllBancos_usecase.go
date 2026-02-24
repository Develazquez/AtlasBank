package application

import (
	"banco-api/Banco/domain/entities"
	"banco-api/Banco/domain/repository"
)

type GetAllBancosUseCase struct {
	repo repository.IBancoRepository
}

func NewGetAllBancosUseCase(repo repository.IBancoRepository) *GetAllBancosUseCase {
	return &GetAllBancosUseCase{repo: repo}
}

func (uc *GetAllBancosUseCase) Execute() ([]*entities.Banco, error) {
	return uc.repo.GetAll()
}
