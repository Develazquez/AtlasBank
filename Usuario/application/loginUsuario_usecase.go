package application

import (
	"banco-api/Usuario/domain/dto"
	"banco-api/Usuario/domain/entities"
	"banco-api/Usuario/domain/repository"
)

type LoginUsuarioUseCase struct {
	repo repository.IUsuarioRepository
}

func NewLoginUsuarioUseCase(repo repository.IUsuarioRepository) *LoginUsuarioUseCase {
	return &LoginUsuarioUseCase{repo: repo}
}

func (uc *LoginUsuarioUseCase) Execute(email, password string) (*entities.Usuario, error) {
	if email == "" || password == "" {
		return nil, entities.ErrCredencialesInvalidas
	}

	usuario, err := uc.repo.LoginUsuario(email, password)
	if err != nil {
		return nil, err
	}
	return usuario, nil
}

func (uc *LoginUsuarioUseCase) ExecuteWithDashboard(email, password string) (*dto.UserDashboardDTO, error) {
	if email == "" || password == "" {
		return nil, entities.ErrCredencialesInvalidas
	}

	dashboard, err := uc.repo.LoginUsuarioWithDashboard(email, password)
	if err != nil {
		return nil, err
	}
	return dashboard, nil
}
