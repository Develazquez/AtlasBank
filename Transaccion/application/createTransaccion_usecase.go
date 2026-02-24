package application

import (
	"banco-api/Transaccion/domain/entities"
	"banco-api/Transaccion/domain/repository"
	"database/sql"
)

type CreateTransaccionUseCase struct {
	repo             repository.ITransaccionRepository
	cuentaRepository CuentaRepository
	db               *sql.DB
}

type CuentaRepository interface {
	GetByID(id int) (*Cuenta, error)
	UpdateSaldo(idCuenta int, nuevoSaldo float64) error
}

type Cuenta struct {
	IDCuenta int
	Saldo    float64
}

func NewCreateTransaccionUseCase(repo repository.ITransaccionRepository, db *sql.DB) *CreateTransaccionUseCase {
	return &CreateTransaccionUseCase{
		repo: repo,
		db:   db,
	}
}

func (uc *CreateTransaccionUseCase) Execute(transaccion *entities.Transaccion) (int, error) {
	if err := transaccion.Validar(); err != nil {
		return 0, err
	}

	// Aquí iría la lógica de incrementar/decrementar saldos según el tipo de transacción
	// Por ahora solo creamos la transacción
	// En una aplicación real, esto sería una transacción de base de datos

	return uc.repo.Create(transaccion)
}
