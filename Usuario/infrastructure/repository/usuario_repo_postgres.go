package repository

import (
	"database/sql"
	"time"

	"banco-api/Usuario/domain/entities"
	"banco-api/Usuario/domain/repository"
)

type UsuarioRepositoryPostgres struct {
	db *sql.DB
}

func NewUsuarioRepositoryPostgres(db *sql.DB) repository.IUsuarioRepository {
	return &UsuarioRepositoryPostgres{db: db}
}

func (r *UsuarioRepositoryPostgres) Create(usuario *entities.Usuario) (int, error) {
	query := `
		INSERT INTO usuarios (nombre, apellido, email, telefono, fecha_nacimiento, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		RETURNING id_usuario
	`

	var id int
	err := r.db.QueryRow(query,
		usuario.Nombre,
		usuario.Apellido,
		usuario.Email,
		usuario.Telefono,
		usuario.FechaNacimiento,
		time.Now(),
		time.Now(),
	).Scan(&id)

	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *UsuarioRepositoryPostgres) GetByID(id int) (*entities.Usuario, error) {
	query := `
		SELECT id_usuario, nombre, apellido, email, telefono, fecha_nacimiento, created_at, updated_at
		FROM usuarios WHERE id_usuario = $1
	`

	usuario := &entities.Usuario{}
	err := r.db.QueryRow(query, id).Scan(
		&usuario.IDUsuario,
		&usuario.Nombre,
		&usuario.Apellido,
		&usuario.Email,
		&usuario.Telefono,
		&usuario.FechaNacimiento,
		&usuario.CreatedAt,
		&usuario.UpdatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return usuario, nil
}

func (r *UsuarioRepositoryPostgres) GetAll() ([]*entities.Usuario, error) {
	query := `
		SELECT id_usuario, nombre, apellido, email, telefono, fecha_nacimiento, created_at, updated_at
		FROM usuarios ORDER BY id_usuario DESC
	`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	usuarios := make([]*entities.Usuario, 0)
	for rows.Next() {
		usuario := &entities.Usuario{}
		err := rows.Scan(
			&usuario.IDUsuario,
			&usuario.Nombre,
			&usuario.Apellido,
			&usuario.Email,
			&usuario.Telefono,
			&usuario.FechaNacimiento,
			&usuario.CreatedAt,
			&usuario.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		usuarios = append(usuarios, usuario)
	}

	return usuarios, nil
}

func (r *UsuarioRepositoryPostgres) Update(usuario *entities.Usuario) error {
	query := `
		UPDATE usuarios
		SET nombre = $1, apellido = $2, email = $3, telefono = $4, fecha_nacimiento = $5, updated_at = $6
		WHERE id_usuario = $7
	`

	_, err := r.db.Exec(query,
		usuario.Nombre,
		usuario.Apellido,
		usuario.Email,
		usuario.Telefono,
		usuario.FechaNacimiento,
		time.Now(),
		usuario.IDUsuario,
	)

	return err
}

func (r *UsuarioRepositoryPostgres) Delete(id int) error {
	query := "DELETE FROM usuarios WHERE id_usuario = $1"
	_, err := r.db.Exec(query, id)
	return err
}

func (r *UsuarioRepositoryPostgres) GetByEmail(email string) (*entities.Usuario, error) {
	query := `
		SELECT id_usuario, nombre, apellido, email, telefono, fecha_nacimiento, created_at, updated_at
		FROM usuarios WHERE email = $1
	`

	usuario := &entities.Usuario{}
	err := r.db.QueryRow(query, email).Scan(
		&usuario.IDUsuario,
		&usuario.Nombre,
		&usuario.Apellido,
		&usuario.Email,
		&usuario.Telefono,
		&usuario.FechaNacimiento,
		&usuario.CreatedAt,
		&usuario.UpdatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return usuario, nil
}
