package application

import (
	"banco-api/Usuario/domain/entities"
	"banco-api/Usuario/domain/repository"
)

type GetUsuarioUseCase struct {
	repo repository.IUsuarioRepository
}

func NewGetUsuarioUseCase(repo repository.IUsuarioRepository) *GetUsuarioUseCase {
	return &GetUsuarioUseCase{repo: repo}
}

func (uc *GetUsuarioUseCase) Execute(id int) (*entities.Usuario, error) {
	usuario, err := uc.repo.GetByID(id)
	if err != nil {
		return nil, err
	}
	if usuario == nil {
		return nil, entities.ErrUsuarioNoEncontrado
	}
	return usuario, nil
}
