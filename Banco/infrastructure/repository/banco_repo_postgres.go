package repository

import (
	"banco-api/Banco/domain/entities"
	"banco-api/Banco/domain/repository"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BancoRepositoryPostgres struct {
	db *gorm.DB
}

func NewBancoRepositoryPostgres(db *gorm.DB) repository.IBancoRepository {
	return &BancoRepositoryPostgres{db: db}
}

func (r *BancoRepositoryPostgres) Create(banco *entities.Banco) (*entities.Banco, error) {
	if banco.ID == uuid.Nil {
		banco.ID = uuid.New()
	}

	if err := r.db.Create(banco).Error; err != nil {
		return nil, err
	}

	return banco, nil
}

func (r *BancoRepositoryPostgres) GetByID(id uuid.UUID) (*entities.Banco, error) {
	var banco entities.Banco
	if err := r.db.Where("id = ?", id).First(&banco).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}

	return &banco, nil
}

func (r *BancoRepositoryPostgres) GetAll() ([]*entities.Banco, error) {
	var bancos []*entities.Banco
	if err := r.db.Order("created_at DESC").Find(&bancos).Error; err != nil {
		return nil, err
	}

	return bancos, nil
}

func (r *BancoRepositoryPostgres) Update(banco *entities.Banco) error {
	return r.db.Save(banco).Error
}

func (r *BancoRepositoryPostgres) Delete(id uuid.UUID) error {
	return r.db.Where("id = ?", id).Delete(&entities.Banco{}).Error
}

func (r *BancoRepositoryPostgres) GetByName(nombre string) (*entities.Banco, error) {
	var banco entities.Banco
	if err := r.db.Where("nombre = ?", nombre).First(&banco).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}

	return &banco, nil
}

func (r *BancoRepositoryPostgres) GetByCodigoSwift(codigo string) (*entities.Banco, error) {
	var banco entities.Banco
	if err := r.db.Where("codigo_swift = ?", codigo).First(&banco).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}

	return &banco, nil
}
