package application

import (
	"banco-api/Banco/domain/entities"
	"banco-api/Banco/domain/repository"
)

type CreateBancoUseCase struct {
	repo repository.IBancoRepository
}

func NewCreateBancoUseCase(repo repository.IBancoRepository) *CreateBancoUseCase {
	return &CreateBancoUseCase{repo: repo}
}

func (uc *CreateBancoUseCase) Execute(banco *entities.Banco) (*entities.Banco, error) {
	if err := banco.Validar(); err != nil {
		return nil, err
	}

	existente, _ := uc.repo.GetByName(banco.Nombre)
	if existente != nil {
		return nil, entities.ErrorBanco{Mensaje: "Ya existe un banco con ese nombre"}
	}

	return uc.repo.Create(banco)
}
