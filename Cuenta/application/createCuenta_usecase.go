package application

import (
	"banco-api/Cuenta/domain/entities"
	"banco-api/Cuenta/domain/repository"

	"github.com/google/uuid"
)

type CreateCuentaUseCase struct {
	repo repository.ICuentaRepository
}

func NewCreateCuentaUseCase(repo repository.ICuentaRepository) *CreateCuentaUseCase {
	return &CreateCuentaUseCase{repo: repo}
}

func (uc *CreateCuentaUseCase) Execute(cuenta *entities.Cuenta) (uuid.UUID, error) {
	if err := cuenta.Validar(); err != nil {
		return uuid.Nil, err
	}

	existente, _ := uc.repo.GetByNumeroCuenta(cuenta.NumeroCuenta)
	if existente != nil {
		return uuid.Nil, entities.ErrNumeroCuentaYaExiste
	}

	return uc.repo.Create(cuenta)
}
