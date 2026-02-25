package application

import (
	"banco-api/Cuenta/domain/entities"
	"banco-api/Cuenta/domain/repository"
)

type UpdateCuentaUseCase struct {
	repo repository.ICuentaRepository
}

func NewUpdateCuentaUseCase(repo repository.ICuentaRepository) *UpdateCuentaUseCase {
	return &UpdateCuentaUseCase{repo: repo}
}

func (uc *UpdateCuentaUseCase) Execute(cuenta *entities.Cuenta) error {
	if err := cuenta.Validar(); err != nil {
		return err
	}

	existente, err := uc.repo.GetByID(cuenta.ID)
	if err != nil || existente == nil {
		return entities.ErrCuentaNoEncontrada
	}

	return uc.repo.Update(cuenta)
}
