package application

import (
	"banco-api/Usuario/domain/dto"
	"banco-api/Usuario/domain/repository"
)

type GetAllUsuariosUseCase struct {
	repo repository.IUsuarioRepository
}

func NewGetAllUsuariosUseCase(repo repository.IUsuarioRepository) *GetAllUsuariosUseCase {
	return &GetAllUsuariosUseCase{repo: repo}
}

func (uc *GetAllUsuariosUseCase) Execute() ([]*dto.UsuarioListDTO, error) {
	return uc.repo.GetAllWithCuentaID()
}
