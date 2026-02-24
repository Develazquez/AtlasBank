package application

import (
	"banco-api/Transaccion/domain/entities"
	"banco-api/Transaccion/domain/repository"
)

type DeleteTransaccionUseCase struct {
	repo repository.ITransaccionRepository
}

func NewDeleteTransaccionUseCase(repo repository.ITransaccionRepository) *DeleteTransaccionUseCase {
	return &DeleteTransaccionUseCase{repo: repo}
}

func (uc *DeleteTransaccionUseCase) Execute(id int) error {
	existente, err := uc.repo.GetByID(id)
	if err != nil || existente == nil {
		return entities.ErrTransaccionNoEncontrada
	}

	return uc.repo.Delete(id)
}
