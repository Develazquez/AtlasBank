package repository

import (
	"database/sql"
	"time"

	"banco-api/Banco/domain/entities"
	"banco-api/Banco/domain/repository"
)

type BancoRepositoryPostgres struct {
	db *sql.DB
}

func NewBancoRepositoryPostgres(db *sql.DB) repository.IBancoRepository {
	return &BancoRepositoryPostgres{db: db}
}

func (r *BancoRepositoryPostgres) Create(banco *entities.Banco) (int, error) {
	query := `
		INSERT INTO banco (nombre, direccion, telefono, email, activo, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		RETURNING id_banco
	`

	var id int
	err := r.db.QueryRow(query,
		banco.Nombre,
		banco.Direccion,
		banco.Telefono,
		banco.Email,
		true,
		time.Now(),
		time.Now(),
	).Scan(&id)

	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *BancoRepositoryPostgres) GetByID(id int) (*entities.Banco, error) {
	query := `
		SELECT id_banco, nombre, direccion, telefono, email, activo, created_at, updated_at
		FROM banco WHERE id_banco = $1
	`

	banco := &entities.Banco{}
	err := r.db.QueryRow(query, id).Scan(
		&banco.IDBanco,
		&banco.Nombre,
		&banco.Direccion,
		&banco.Telefono,
		&banco.Email,
		&banco.Activo,
		&banco.CreatedAt,
		&banco.UpdatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return banco, nil
}

func (r *BancoRepositoryPostgres) GetAll() ([]*entities.Banco, error) {
	query := `
		SELECT id_banco, nombre, direccion, telefono, email, activo, created_at, updated_at
		FROM banco ORDER BY id_banco DESC
	`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	bancos := make([]*entities.Banco, 0)
	for rows.Next() {
		banco := &entities.Banco{}
		err := rows.Scan(
			&banco.IDBanco,
			&banco.Nombre,
			&banco.Direccion,
			&banco.Telefono,
			&banco.Email,
			&banco.Activo,
			&banco.CreatedAt,
			&banco.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		bancos = append(bancos, banco)
	}

	return bancos, nil
}

func (r *BancoRepositoryPostgres) Update(banco *entities.Banco) error {
	query := `
		UPDATE banco 
		SET nombre = $1, direccion = $2, telefono = $3, email = $4, activo = $5, updated_at = $6
		WHERE id_banco = $7
	`

	_, err := r.db.Exec(query,
		banco.Nombre,
		banco.Direccion,
		banco.Telefono,
		banco.Email,
		banco.Activo,
		time.Now(),
		banco.IDBanco,
	)

	return err
}

func (r *BancoRepositoryPostgres) Delete(id int) error {
	query := "DELETE FROM banco WHERE id_banco = $1"
	_, err := r.db.Exec(query, id)
	return err
}

func (r *BancoRepositoryPostgres) GetByName(nombre string) (*entities.Banco, error) {
	query := `
		SELECT id_banco, nombre, direccion, telefono, email, activo, created_at, updated_at
		FROM banco WHERE nombre = $1
	`

	banco := &entities.Banco{}
	err := r.db.QueryRow(query, nombre).Scan(
		&banco.IDBanco,
		&banco.Nombre,
		&banco.Direccion,
		&banco.Telefono,
		&banco.Email,
		&banco.Activo,
		&banco.CreatedAt,
		&banco.UpdatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return banco, nil
}
