package application

import (
	"banco-api/Cuenta/domain/entities"
	"banco-api/Cuenta/domain/repository"
)

type GetAllCuentasUseCase struct {
	repo repository.ICuentaRepository
}

func NewGetAllCuentasUseCase(repo repository.ICuentaRepository) *GetAllCuentasUseCase {
	return &GetAllCuentasUseCase{repo: repo}
}

func (uc *GetAllCuentasUseCase) Execute() ([]*entities.Cuenta, error) {
	return uc.repo.GetAll()
}
