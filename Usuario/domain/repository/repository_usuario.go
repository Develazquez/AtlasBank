package repository

import "banco-api/Usuario/domain/entities"

type IUsuarioRepository interface {
	Create(usuario *entities.Usuario) (int, error)
	GetByID(id int) (*entities.Usuario, error)
	GetAll() ([]*entities.Usuario, error)
	Update(usuario *entities.Usuario) error
	Delete(id int) error
	GetByEmail(email string) (*entities.Usuario, error)
}
