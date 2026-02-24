package application

import (
	"banco-api/Cuenta/domain/entities"
	"banco-api/Cuenta/domain/repository"
)

type CreateCuentaUseCase struct {
	repo repository.ICuentaRepository
}

func NewCreateCuentaUseCase(repo repository.ICuentaRepository) *CreateCuentaUseCase {
	return &CreateCuentaUseCase{repo: repo}
}

func (uc *CreateCuentaUseCase) Execute(cuenta *entities.Cuenta) (int, error) {
	if err := cuenta.Validar(); err != nil {
		return 0, err
	}

	existente, _ := uc.repo.GetByNumeroCuenta(cuenta.NumeroCuenta)
	if existente != nil {
		return 0, entities.ErrNumeroCuentaYaExiste
	}

	return uc.repo.Create(cuenta)
}
