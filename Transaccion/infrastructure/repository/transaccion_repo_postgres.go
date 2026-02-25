package repository

import (
	"github.com/google/uuid"
	"gorm.io/gorm"

	"banco-api/Transaccion/domain/entities"
	"banco-api/Transaccion/domain/repository"
)

type TransaccionRepositoryPostgres struct {
	db *gorm.DB
}

func NewTransaccionRepositoryPostgres(db *gorm.DB) repository.ITransaccionRepository {
	return &TransaccionRepositoryPostgres{db: db}
}

func (r *TransaccionRepositoryPostgres) Create(transaccion *entities.Transaccion) (uuid.UUID, error) {
	if transaccion.ID == uuid.Nil {
		transaccion.ID = uuid.New()
	}
	result := r.db.Create(transaccion)
	return transaccion.ID, result.Error
}

func (r *TransaccionRepositoryPostgres) GetByID(id uuid.UUID) (*entities.Transaccion, error) {
	transaccion := &entities.Transaccion{}
	result := r.db.First(transaccion, "id = ?", id)

	if result.Error == gorm.ErrRecordNotFound {
		return nil, nil
	}
	if result.Error != nil {
		return nil, result.Error
	}

	return transaccion, nil
}

func (r *TransaccionRepositoryPostgres) GetAll() ([]*entities.Transaccion, error) {
	var transacciones []*entities.Transaccion
	result := r.db.Order("created_at DESC").Find(&transacciones)
	return transacciones, result.Error
}

func (r *TransaccionRepositoryPostgres) Delete(id uuid.UUID) error {
	result := r.db.Delete(&entities.Transaccion{}, "id = ?", id)
	return result.Error
}

func (r *TransaccionRepositoryPostgres) GetTransactionsByCuenta(idCuenta uuid.UUID) ([]*entities.Transaccion, error) {
	var transacciones []*entities.Transaccion
	result := r.db.Where("cuenta_origen_id = ? OR cuenta_destino_id = ?", idCuenta, idCuenta).
		Order("created_at DESC").
		Find(&transacciones)
	return transacciones, result.Error
}
