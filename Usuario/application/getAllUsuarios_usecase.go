package application

import (
	"banco-api/Usuario/domain/entities"
	"banco-api/Usuario/domain/repository"
)

type GetAllUsuariosUseCase struct {
	repo repository.IUsuarioRepository
}

func NewGetAllUsuariosUseCase(repo repository.IUsuarioRepository) *GetAllUsuariosUseCase {
	return &GetAllUsuariosUseCase{repo: repo}
}

func (uc *GetAllUsuariosUseCase) Execute() ([]*entities.Usuario, error) {
	return uc.repo.GetAll()
}
