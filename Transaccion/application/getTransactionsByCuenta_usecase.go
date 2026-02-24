package application

import (
	"banco-api/Transaccion/domain/entities"
	"banco-api/Transaccion/domain/repository"
)

type GetTransactionsByCuentaUseCase struct {
	repo repository.ITransaccionRepository
}

func NewGetTransactionsByCuentaUseCase(repo repository.ITransaccionRepository) *GetTransactionsByCuentaUseCase {
	return &GetTransactionsByCuentaUseCase{repo: repo}
}

func (uc *GetTransactionsByCuentaUseCase) Execute(idCuenta int) ([]*entities.Transaccion, error) {
	return uc.repo.GetTransactionsByCuenta(idCuenta)
}
