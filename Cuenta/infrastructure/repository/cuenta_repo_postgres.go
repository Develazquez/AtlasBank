package repository

import (
	"database/sql"
	"time"

	"banco-api/Cuenta/domain/entities"
	"banco-api/Cuenta/domain/repository"
)

type CuentaRepositoryPostgres struct {
	db *sql.DB
}

func NewCuentaRepositoryPostgres(db *sql.DB) repository.ICuentaRepository {
	return &CuentaRepositoryPostgres{db: db}
}

func (r *CuentaRepositoryPostgres) Create(cuenta *entities.Cuenta) (int, error) {
	query := `
		INSERT INTO cuenta (numero_cuenta, tipo_cuenta, saldo, id_usuario, id_banco, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		RETURNING id_cuenta
	`

	var id int
	err := r.db.QueryRow(query,
		cuenta.NumeroCuenta,
		cuenta.TipoCuenta,
		cuenta.Saldo,
		cuenta.IDUsuario,
		cuenta.IDBanco,
		time.Now(),
		time.Now(),
	).Scan(&id)

	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *CuentaRepositoryPostgres) GetByID(id int) (*entities.Cuenta, error) {
	query := `
		SELECT id_cuenta, numero_cuenta, tipo_cuenta, saldo, id_usuario, id_banco, created_at, updated_at
		FROM cuenta WHERE id_cuenta = $1
	`

	cuenta := &entities.Cuenta{}
	err := r.db.QueryRow(query, id).Scan(
		&cuenta.IDCuenta,
		&cuenta.NumeroCuenta,
		&cuenta.TipoCuenta,
		&cuenta.Saldo,
		&cuenta.IDUsuario,
		&cuenta.IDBanco,
		&cuenta.CreatedAt,
		&cuenta.UpdatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return cuenta, nil
}

func (r *CuentaRepositoryPostgres) GetAll() ([]*entities.Cuenta, error) {
	query := `
		SELECT id_cuenta, numero_cuenta, tipo_cuenta, saldo, id_usuario, id_banco, created_at, updated_at
		FROM cuenta ORDER BY id_cuenta DESC
	`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	cuentas := make([]*entities.Cuenta, 0)
	for rows.Next() {
		cuenta := &entities.Cuenta{}
		err := rows.Scan(
			&cuenta.IDCuenta,
			&cuenta.NumeroCuenta,
			&cuenta.TipoCuenta,
			&cuenta.Saldo,
			&cuenta.IDUsuario,
			&cuenta.IDBanco,
			&cuenta.CreatedAt,
			&cuenta.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		cuentas = append(cuentas, cuenta)
	}

	return cuentas, nil
}

func (r *CuentaRepositoryPostgres) Update(cuenta *entities.Cuenta) error {
	query := `
		UPDATE cuenta
		SET numero_cuenta = $1, tipo_cuenta = $2, saldo = $3, id_usuario = $4, id_banco = $5, updated_at = $6
		WHERE id_cuenta = $7
	`

	_, err := r.db.Exec(query,
		cuenta.NumeroCuenta,
		cuenta.TipoCuenta,
		cuenta.Saldo,
		cuenta.IDUsuario,
		cuenta.IDBanco,
		time.Now(),
		cuenta.IDCuenta,
	)

	return err
}

func (r *CuentaRepositoryPostgres) Delete(id int) error {
	query := "DELETE FROM cuenta WHERE id_cuenta = $1"
	_, err := r.db.Exec(query, id)
	return err
}

func (r *CuentaRepositoryPostgres) GetByNumeroCuenta(numero string) (*entities.Cuenta, error) {
	query := `
		SELECT id_cuenta, numero_cuenta, tipo_cuenta, saldo, id_usuario, id_banco, created_at, updated_at
		FROM cuenta WHERE numero_cuenta = $1
	`

	cuenta := &entities.Cuenta{}
	err := r.db.QueryRow(query, numero).Scan(
		&cuenta.IDCuenta,
		&cuenta.NumeroCuenta,
		&cuenta.TipoCuenta,
		&cuenta.Saldo,
		&cuenta.IDUsuario,
		&cuenta.IDBanco,
		&cuenta.CreatedAt,
		&cuenta.UpdatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return cuenta, nil
}

func (r *CuentaRepositoryPostgres) GetCuentasByUsuario(idUsuario int) ([]*entities.Cuenta, error) {
	query := `
		SELECT id_cuenta, numero_cuenta, tipo_cuenta, saldo, id_usuario, id_banco, created_at, updated_at
		FROM cuenta WHERE id_usuario = $1 ORDER BY id_cuenta DESC
	`

	rows, err := r.db.Query(query, idUsuario)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	cuentas := make([]*entities.Cuenta, 0)
	for rows.Next() {
		cuenta := &entities.Cuenta{}
		err := rows.Scan(
			&cuenta.IDCuenta,
			&cuenta.NumeroCuenta,
			&cuenta.TipoCuenta,
			&cuenta.Saldo,
			&cuenta.IDUsuario,
			&cuenta.IDBanco,
			&cuenta.CreatedAt,
			&cuenta.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		cuentas = append(cuentas, cuenta)
	}

	return cuentas, nil
}
