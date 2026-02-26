package application

import (
	"banco-api/Cuenta/domain/repository"
	"banco-api/Transaccion/domain/entities"
	transaccionRepo "banco-api/Transaccion/domain/repository"
	cuentaEntities "banco-api/Cuenta/domain/entities"


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

	err := uc.db.Transaction(func(tx *gorm.DB) error {

		// Descontar saldo de cuenta origen
		if transaccion.CuentaOrigenID != nil {
			var origen cuentaEntities.Cuenta
			if err := tx.First(&origen, "id = ?", transaccion.CuentaOrigenID).Error; err != nil {
				return entities.ErrCuentaOrigen
			}
			if origen.Saldo < transaccion.Monto {
				return entities.ErrSaldoInsuficiente
			}
			origen.Saldo -= transaccion.Monto
			if err := tx.Save(&origen).Error; err != nil {
				return err
			}
		}

		// Acreditar saldo a cuenta destino
		if transaccion.CuentaDestinoID != nil {
			var destino cuentaEntities.Cuenta
			if err := tx.First(&destino, "id = ?", transaccion.CuentaDestinoID).Error; err != nil {
				return entities.ErrCuentaDestino
			}
			destino.Saldo += transaccion.Monto
			if err := tx.Save(&destino).Error; err != nil {
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