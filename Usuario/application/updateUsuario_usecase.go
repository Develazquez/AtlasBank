package application

import (
	"banco-api/Usuario/domain/entities"
	"banco-api/Usuario/domain/repository"
)

type UpdateUsuarioUseCase struct {
	repo repository.IUsuarioRepository
}

func NewUpdateUsuarioUseCase(repo repository.IUsuarioRepository) *UpdateUsuarioUseCase {
	return &UpdateUsuarioUseCase{repo: repo}
}

func (uc *UpdateUsuarioUseCase) Execute(usuario *entities.Usuario) error {
	if err := usuario.Validar(); err != nil {
		return err
	}

	existente, err := uc.repo.GetByID(usuario.ID)
	if err != nil || existente == nil {
		return entities.ErrUsuarioNoEncontrado
	}

	return uc.repo.Update(usuario)
}
