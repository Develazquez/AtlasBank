package repository

import (
	"github.com/google/uuid"
	"gorm.io/gorm"

	"banco-api/Cuenta/domain/entities"
	"banco-api/Cuenta/domain/repository"
)

type CuentaRepositoryPostgres struct {
	db *gorm.DB
}

func NewCuentaRepositoryPostgres(db *gorm.DB) repository.ICuentaRepository {
	return &CuentaRepositoryPostgres{db: db}
}

func (r *CuentaRepositoryPostgres) Create(cuenta *entities.Cuenta) (uuid.UUID, error) {
	if cuenta.ID == uuid.Nil {
		cuenta.ID = uuid.New()
	}
	result := r.db.Create(cuenta)
	return cuenta.ID, result.Error
}

func (r *CuentaRepositoryPostgres) GetByID(id uuid.UUID) (*entities.Cuenta, error) {
	cuenta := &entities.Cuenta{}
	result := r.db.First(cuenta, "id = ?", id)

	if result.Error == gorm.ErrRecordNotFound {
		return nil, nil
	}
	if result.Error != nil {
		return nil, result.Error
	}

	return cuenta, nil
}

func (r *CuentaRepositoryPostgres) GetAll() ([]*entities.Cuenta, error) {
	var cuentas []*entities.Cuenta
	result := r.db.Order("created_at DESC").Find(&cuentas)
	return cuentas, result.Error
}

func (r *CuentaRepositoryPostgres) Update(cuenta *entities.Cuenta) error {
	result := r.db.Save(cuenta)
	return result.Error
}

func (r *CuentaRepositoryPostgres) Delete(id uuid.UUID) error {
	result := r.db.Delete(&entities.Cuenta{}, "id = ?", id)
	return result.Error
}

func (r *CuentaRepositoryPostgres) GetByNumeroCuenta(numero string) (*entities.Cuenta, error) {
	cuenta := &entities.Cuenta{}
	result := r.db.First(cuenta, "numero_cuenta = ?", numero)

	if result.Error == gorm.ErrRecordNotFound {
		return nil, nil
	}
	if result.Error != nil {
		return nil, result.Error
	}

	return cuenta, nil
}

func (r *CuentaRepositoryPostgres) GetCuentasByUsuario(idUsuario uuid.UUID) ([]*entities.Cuenta, error) {
	var cuentas []*entities.Cuenta
	result := r.db.Where("usuario_id = ?", idUsuario).Order("created_at DESC").Find(&cuentas)
	return cuentas, result.Error
}
