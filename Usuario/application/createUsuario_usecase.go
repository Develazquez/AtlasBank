package application

import (
	"banco-api/Usuario/domain/entities"
	"banco-api/Usuario/domain/repository"

	"github.com/google/uuid"
)

type CreateUsuarioUseCase struct {
	repo repository.IUsuarioRepository
}

func NewCreateUsuarioUseCase(repo repository.IUsuarioRepository) *CreateUsuarioUseCase {
	return &CreateUsuarioUseCase{repo: repo}
}

func (uc *CreateUsuarioUseCase) Execute(usuario *entities.Usuario) (uuid.UUID, error) {
	if err := usuario.Validar(); err != nil {
		return uuid.Nil, err
	}

	existente, _ := uc.repo.GetByEmail(usuario.Email)
	if existente != nil {
		return uuid.Nil, entities.ErrEmailYaExiste
	}

	return uc.repo.Create(usuario)
}
