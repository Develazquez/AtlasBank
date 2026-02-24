package application

import (
	"banco-api/Transaccion/domain/entities"
	"banco-api/Transaccion/domain/repository"
)

type GetAllTransaccionesUseCase struct {
	repo repository.ITransaccionRepository
}

func NewGetAllTransaccionesUseCase(repo repository.ITransaccionRepository) *GetAllTransaccionesUseCase {
	return &GetAllTransaccionesUseCase{repo: repo}
}

func (uc *GetAllTransaccionesUseCase) Execute() ([]*entities.Transaccion, error) {
	return uc.repo.GetAll()
}
