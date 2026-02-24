package application

import (
	"banco-api/Usuario/domain/entities"
	"banco-api/Usuario/domain/repository"
)

var UsuarioRepo repository.IUsuarioRepository

func LoginUsuarioUsecase(email, password string) (*entities.Usuario, error) {
	if UsuarioRepo == nil {
		return nil, entities.ErrUsuarioNoEncontrado
	}
	usuario, err := UsuarioRepo.LoginUsuario(email, password)
	if err != nil {
		return nil, err
	}
	return usuario, nil
}
