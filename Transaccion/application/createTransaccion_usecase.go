package application

import (
	"banco-api/Cuenta/domain/repository"
	"banco-api/Transaccion/domain/entities"
	transaccionRepo "banco-api/Transaccion/domain/repository"
	cuentaEntities "banco-api/Cuenta/domain/entities"


	"github.com/google/uuid"
	"gorm.io/gorm"
	"fmt"
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

	fmt.Println(">>> Iniciando transacción BD")

	err := uc.db.Transaction(func(tx *gorm.DB) error {

		if transaccion.CuentaOrigenID != nil {
			fmt.Printf(">>> Buscando cuenta origen: %s\n", transaccion.CuentaOrigenID.String())
			var origen cuentaEntities.Cuenta
			if err := tx.First(&origen, "id = ?", transaccion.CuentaOrigenID).Error; err != nil {
				fmt.Printf("❌ Error buscando origen: %s\n", err.Error())
				return entities.ErrCuentaOrigen
			}
			fmt.Printf(">>> Saldo origen antes: %f\n", origen.Saldo)
			if origen.Saldo < transaccion.Monto {
				fmt.Println("❌ Saldo insuficiente")
				return entities.ErrSaldoInsuficiente
			}
			origen.Saldo -= transaccion.Monto
			fmt.Printf(">>> Saldo origen después: %f\n", origen.Saldo)
			if err := tx.Save(&origen).Error; err != nil {
				fmt.Printf("❌ Error guardando origen: %s\n", err.Error())
				return err
			}
		}

		if transaccion.CuentaDestinoID != nil {
			fmt.Printf(">>> Buscando cuenta destino: %s\n", transaccion.CuentaDestinoID.String())
			var destino cuentaEntities.Cuenta
			if err := tx.First(&destino, "id = ?", transaccion.CuentaDestinoID).Error; err != nil {
				fmt.Printf("❌ Error buscando destino: %s\n", err.Error())
				return entities.ErrCuentaDestino
			}
			fmt.Printf(">>> Saldo destino antes: %f\n", destino.Saldo)
			destino.Saldo += transaccion.Monto
			fmt.Printf(">>> Saldo destino después: %f\n", destino.Saldo)
			if err := tx.Save(&destino).Error; err != nil {
				fmt.Printf("❌ Error guardando destino: %s\n", err.Error())
				return err
			}
		}

		fmt.Println(">>> Transacción BD completada")
		return nil
	})

	if err != nil {
		fmt.Printf("❌ Transacción BD falló: %s\n", err.Error())
		return uuid.Nil, err
	}

	return uc.repo.Create(transaccion)
}