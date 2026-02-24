package repository

import "banco-api/Transaccion/domain/entities"

type ITransaccionRepository interface {
	Create(transaccion *entities.Transaccion) (int, error)
	GetByID(id int) (*entities.Transaccion, error)
	GetAll() ([]*entities.Transaccion, error)
	Delete(id int) error
	GetTransactionsByCuenta(idCuenta int) ([]*entities.Transaccion, error)
}
