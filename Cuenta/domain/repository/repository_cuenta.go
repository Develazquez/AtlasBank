package repository

import "banco-api/Cuenta/domain/entities"

type ICuentaRepository interface {
	Create(cuenta *entities.Cuenta) (int, error)
	GetByID(id int) (*entities.Cuenta, error)
	GetAll() ([]*entities.Cuenta, error)
	Update(cuenta *entities.Cuenta) error
	Delete(id int) error
	GetByNumeroCuenta(numero string) (*entities.Cuenta, error)
	GetCuentasByUsuario(idUsuario int) ([]*entities.Cuenta, error)
}
