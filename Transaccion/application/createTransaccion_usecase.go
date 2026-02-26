package application

import (
	"banco-api/Cuenta/domain/repository"
	"banco-api/Transaccion/domain/entities"
	transaccionRepo "banco-api/Transaccion/domain/repository"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CreateTransaccionUseCase struct {
	repo       transaccionRepo.ITransaccionRepository
	cuentaRepo repository.ICuentaRepository
	db         *gorm.DB
}

func NewCreateTransaccionUseCase(
	repo transaccionRepo.ITransaccionRepository,
	cuentaRepo repository.ICuentaRepository,
	db *gorm.DB,
) *CreateTransaccionUseCase {
	return &CreateTransaccionUseCase{
		repo:       repo,
		cuentaRepo: cuentaRepo,
		db:         db,
	}
}

func (uc *CreateTransaccionUseCase) Execute(transaccion *entities.Transaccion) (uuid.UUID, error) {
	if err := transaccion.Validar(); err != nil {
		return uuid.Nil, err
	}

	// Actualizar saldos dentro de una transacción BD
	err := uc.db.Transaction(func(tx *gorm.DB) error {

		// Descontar saldo de cuenta origen
		if transaccion.CuentaOrigenID != nil {
			origen, err := uc.cuentaRepo.GetByID(*transaccion.CuentaOrigenID)
			if err != nil || origen == nil {
				return entities.ErrCuentaOrigen
			}
			if origen.Saldo < transaccion.Monto {
				return entities.ErrSaldoInsuficiente
			}
			origen.Saldo -= transaccion.Monto
			if err := uc.cuentaRepo.Update(origen); err != nil {
				return err
			}
		}

		// Acreditar saldo a cuenta destino
		if transaccion.CuentaDestinoID != nil {
			destino, err := uc.cuentaRepo.GetByID(*transaccion.CuentaDestinoID)
			if err != nil || destino == nil {
				return entities.ErrCuentaDestino
			}
			destino.Saldo += transaccion.Monto
			if err := uc.cuentaRepo.Update(destino); err != nil {
				return err
			}
		}

		return nil
	})

	if err != nil {
		return uuid.Nil, err
	}

	return uc.repo.Create(transaccion)
}