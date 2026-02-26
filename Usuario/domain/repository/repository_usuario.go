package repository

import (
	"banco-api/Usuario/domain/dto"
	"banco-api/Usuario/domain/entities"

	"github.com/google/uuid"
)

type IUsuarioRepository interface {
	Create(usuario *entities.Usuario) (uuid.UUID, error)
	GetByID(id uuid.UUID) (*entities.Usuario, error)
	GetAll() ([]*entities.Usuario, error)
	Update(usuario *entities.Usuario) error
	Delete(id uuid.UUID) error
	GetByEmail(email string) (*entities.Usuario, error)
	LoginUsuario(email, password string) (*entities.Usuario, error)
	LoginUsuarioWithDashboard(email, password string) (*dto.UserDashboardDTO, error)
}
