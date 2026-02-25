package application

import (
	"banco-api/Cuenta/domain/entities"
	"banco-api/Cuenta/domain/repository"

	"github.com/google/uuid"
)

type GetCuentasByUsuarioUseCase struct {
	repo repository.ICuentaRepository
}

func NewGetCuentasByUsuarioUseCase(repo repository.ICuentaRepository) *GetCuentasByUsuarioUseCase {
	return &GetCuentasByUsuarioUseCase{repo: repo}
}

func (uc *GetCuentasByUsuarioUseCase) Execute(idUsuario uuid.UUID) ([]*entities.Cuenta, error) {
	return uc.repo.GetCuentasByUsuario(idUsuario)
}
