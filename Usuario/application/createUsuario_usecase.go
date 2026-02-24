package application

import (
	"banco-api/Usuario/domain/entities"
	"banco-api/Usuario/domain/repository"
)

type CreateUsuarioUseCase struct {
	repo repository.IUsuarioRepository
}

func NewCreateUsuarioUseCase(repo repository.IUsuarioRepository) *CreateUsuarioUseCase {
	return &CreateUsuarioUseCase{repo: repo}
}

func (uc *CreateUsuarioUseCase) Execute(usuario *entities.Usuario) (int, error) {
	if err := usuario.Validar(); err != nil {
		return 0, err
	}

	existente, _ := uc.repo.GetByEmail(usuario.Email)
	if existente != nil {
		return 0, entities.ErrEmailYaExiste
	}

	return uc.repo.Create(usuario)
}
