package repository

import "banco-api/Banco/domain/entities"

type IBancoRepository interface {

	Create(banco *entities.Banco) (int, error)
	GetByID(id int) (*entities.Banco, error)
	GetAll() ([]*entities.Banco, error)
	Update(banco *entities.Banco) error
	Delete(id int) error
	GetByName(nombre string) (*entities.Banco, error)
}
