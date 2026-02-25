package application

import (
	"banco-api/Transaccion/domain/entities"
	"banco-api/Transaccion/domain/repository"

	"github.com/google/uuid"
)

type GetTransactionsByCuentaUseCase struct {
	repo repository.ITransaccionRepository
}

func NewGetTransactionsByCuentaUseCase(repo repository.ITransaccionRepository) *GetTransactionsByCuentaUseCase {
	return &GetTransactionsByCuentaUseCase{repo: repo}
}

func (uc *GetTransactionsByCuentaUseCase) Execute(idCuenta uuid.UUID) ([]*entities.Transaccion, error) {
	return uc.repo.GetTransactionsByCuenta(idCuenta)
}
