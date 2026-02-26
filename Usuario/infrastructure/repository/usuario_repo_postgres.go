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
	// Estructura auxiliar para hacer el JOIN
	var result struct {
		UsuarioID             uuid.UUID
		Nombre                string
		ApellidoPaterno       string
		ApellidoMaterno       string
		Email                 string
		PasswordHash          string
		CuentaID              uuid.UUID
		Saldo                 float64
		TipoTarjeta           string
		UltimosDigitosTarjeta string
		FechaExpiracion       string
		NombreTarjeta         string
		NumeroCuenta          string
		CLABE                 string
		IBAN                  string
		BancoID               uuid.UUID
		CodigoSwift           string
		BancoNombre           string
	}

	// Query que hace JOIN entre usuarios, cuenta y banco
	query := r.db.
		Select(
			"u.id as usuario_id",
			"u.nombre as nombre",
			"u.apellido_paterno as apellido_paterno",
			"u.apellido_materno as apellido_materno",
			"u.email as email",
			"u.password_hash as password_hash",
			"c.id as cuenta_id",
			"c.saldo as saldo",
			"c.tipo_tarjeta as tipo_tarjeta",
			"c.ultimos_digitos_tarjeta as ultimos_digitos_tarjeta",
			"c.fecha_expiracion as fecha_expiracion",
			"c.nombre_tarjeta as nombre_tarjeta",
			"c.numero_cuenta as numero_cuenta",
			"c.clabe as clabe",
			"c.iban as iban",
			"b.id as banco_id",
			"b.codigo_swift as codigo_swift",
			"b.nombre as banco_nombre",
		).
		Table("usuarios u").
		Joins("JOIN cuenta c ON c.usuario_id = u.id").
		Joins("JOIN banco b ON b.id = u.banco_id").
		Where("u.email = ? AND u.activo = ? AND c.estado = ?", email, true, "ACTIVA").
		First(&result)

	if query.Error == gorm.ErrRecordNotFound {
		return nil, entities.ErrUsuarioNoEncontrado
	}
	if query.Error != nil {
		return nil, query.Error
	}

	// Validar contraseña
	if result.PasswordHash != password {
		return nil, entities.ErrUsuarioNoEncontrado
	}

	// Construir DTO con el nombre completo
	nombreCompleto := result.Nombre + " " + result.ApellidoPaterno
	if result.ApellidoMaterno != "" {
		nombreCompleto += " " + result.ApellidoMaterno
	}

	dashboard := &dto.UserDashboardDTO{
		UsuarioID:      result.UsuarioID,
		NombreCompleto: nombreCompleto,
		Email:          result.Email,
		Wallet:         result.Saldo,
		CuentaID:       result.CuentaID,
		NumeroCuenta:   result.NumeroCuenta,
		CLABE:          result.CLABE,
		IBAN:           result.IBAN,
		CodigoSwift:    result.CodigoSwift,
		BancoNombre:    result.BancoNombre,
		Card: dto.CardDTO{
			NameCard:   result.TipoTarjeta,
			NumCard:    result.UltimosDigitosTarjeta,
			Expires:    result.FechaExpiracion,
			CardHolder: result.NombreTarjeta,
		},
	}

	return dashboard, nil
}
