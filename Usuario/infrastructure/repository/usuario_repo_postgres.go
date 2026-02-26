package repository

import (
	"github.com/google/uuid"
	"gorm.io/gorm"

	"banco-api/Usuario/domain/dto"
	"banco-api/Usuario/domain/entities"
	"banco-api/Usuario/domain/repository"
)

type UsuarioRepositoryPostgres struct {
	db *gorm.DB
}

func NewUsuarioRepositoryPostgres(db *gorm.DB) repository.IUsuarioRepository {
	return &UsuarioRepositoryPostgres{db: db}
}

func (r *UsuarioRepositoryPostgres) Create(usuario *entities.Usuario) (uuid.UUID, error) {
	if usuario.ID == uuid.Nil {
		usuario.ID = uuid.New()
	}
	result := r.db.Create(usuario)
	return usuario.ID, result.Error
}

func (r *UsuarioRepositoryPostgres) GetByID(id uuid.UUID) (*entities.Usuario, error) {
	usuario := &entities.Usuario{}
	result := r.db.First(usuario, "id = ?", id)

	if result.Error == gorm.ErrRecordNotFound {
		return nil, nil
	}
	if result.Error != nil {
		return nil, result.Error
	}

	return usuario, nil
}

func (r *UsuarioRepositoryPostgres) GetAll() ([]*entities.Usuario, error) {
	var usuarios []*entities.Usuario
	result := r.db.Order("created_at DESC").Find(&usuarios)
	return usuarios, result.Error
}

func (r *UsuarioRepositoryPostgres) Update(usuario *entities.Usuario) error {
	result := r.db.Save(usuario)
	return result.Error
}

func (r *UsuarioRepositoryPostgres) Delete(id uuid.UUID) error {
	result := r.db.Delete(&entities.Usuario{}, "id = ?", id)
	return result.Error
}

func (r *UsuarioRepositoryPostgres) GetByEmail(email string) (*entities.Usuario, error) {
	usuario := &entities.Usuario{}
	result := r.db.First(usuario, "email = ?", email)

	if result.Error == gorm.ErrRecordNotFound {
		return nil, nil
	}
	if result.Error != nil {
		return nil, result.Error
	}

	return usuario, nil
}

func (r *UsuarioRepositoryPostgres) LoginUsuario(email, password string) (*entities.Usuario, error) {
	usuario := &entities.Usuario{}
	result := r.db.First(usuario, "email = ?", email)

	if result.Error == gorm.ErrRecordNotFound {
		return nil, entities.ErrUsuarioNoEncontrado
	}
	if result.Error != nil {
		return nil, result.Error
	}

	if usuario.PasswordHash != password {
		return nil, entities.ErrUsuarioNoEncontrado
	}
	return usuario, nil
}

func (r *UsuarioRepositoryPostgres) LoginUsuarioWithDashboard(email, password string) (*dto.UserDashboardDTO, error) {
	usuario := &entities.Usuario{}
	if err := r.db.First(usuario, "email = ? AND activo = ?", email, true).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, entities.ErrUsuarioNoEncontrado
		}
		return nil, err
	}

	if usuario.PasswordHash != password {
		return nil, entities.ErrCredencialesInvalidas
	}

	// Estructura auxiliar para mapear la vista
	var dashboard dto.UserDashboardDTO

	sqlQuery := `
		SELECT 
			id,
			name,
			wallet,
			card,
			recently_inf
		FROM v_usuario_dashboard
		WHERE id = $1
		LIMIT 1
	`

	query := r.db.Raw(sqlQuery, usuario.ID.String()).Scan(&dashboard)

	if query.Error != nil {
		if query.Error == gorm.ErrRecordNotFound {
			return nil, entities.ErrUsuarioNoEncontrado
		}
		return nil, query.Error
	}

	return &dashboard, nil
}
