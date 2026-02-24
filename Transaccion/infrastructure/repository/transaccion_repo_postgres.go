package repository

import (
	"database/sql"
	"time"

	"banco-api/Transaccion/domain/entities"
	"banco-api/Transaccion/domain/repository"
)

type TransaccionRepositoryPostgres struct {
	db *sql.DB
}

func NewTransaccionRepositoryPostgres(db *sql.DB) repository.ITransaccionRepository {
	return &TransaccionRepositoryPostgres{db: db}
}

func (r *TransaccionRepositoryPostgres) Create(transaccion *entities.Transaccion) (int, error) {
	query := `
		INSERT INTO transacciones (tipo_transaccion, monto, fecha, cuenta_origen, cuenta_destino, descripcion)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id_transaccion
	`

	var id int
	err := r.db.QueryRow(query,
		transaccion.TipoTransaccion,
		transaccion.Monto,
		time.Now(),
		transaccion.CuentaOrigen,
		transaccion.CuentaDestino,
		transaccion.Descripcion,
	).Scan(&id)

	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *TransaccionRepositoryPostgres) GetByID(id int) (*entities.Transaccion, error) {
	query := `
		SELECT id_transaccion, tipo_transaccion, monto, fecha, cuenta_origen, cuenta_destino, descripcion
		FROM transacciones WHERE id_transaccion = $1
	`

	transaccion := &entities.Transaccion{}
	err := r.db.QueryRow(query, id).Scan(
		&transaccion.IDTransaccion,
		&transaccion.TipoTransaccion,
		&transaccion.Monto,
		&transaccion.Fecha,
		&transaccion.CuentaOrigen,
		&transaccion.CuentaDestino,
		&transaccion.Descripcion,
	)

	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return transaccion, nil
}

func (r *TransaccionRepositoryPostgres) GetAll() ([]*entities.Transaccion, error) {
	query := `
		SELECT id_transaccion, tipo_transaccion, monto, fecha, cuenta_origen, cuenta_destino, descripcion
		FROM transacciones ORDER BY id_transaccion DESC
	`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	transacciones := make([]*entities.Transaccion, 0)
	for rows.Next() {
		transaccion := &entities.Transaccion{}
		err := rows.Scan(
			&transaccion.IDTransaccion,
			&transaccion.TipoTransaccion,
			&transaccion.Monto,
			&transaccion.Fecha,
			&transaccion.CuentaOrigen,
			&transaccion.CuentaDestino,
			&transaccion.Descripcion,
		)
		if err != nil {
			return nil, err
		}
		transacciones = append(transacciones, transaccion)
	}

	return transacciones, nil
}

func (r *TransaccionRepositoryPostgres) Delete(id int) error {
	query := "DELETE FROM transacciones WHERE id_transaccion = $1"
	_, err := r.db.Exec(query, id)
	return err
}

func (r *TransaccionRepositoryPostgres) GetTransactionsByCuenta(idCuenta int) ([]*entities.Transaccion, error) {
	query := `
		SELECT id_transaccion, tipo_transaccion, monto, fecha, cuenta_origen, cuenta_destino, descripcion
		FROM transacciones 
		WHERE cuenta_origen = $1 OR cuenta_destino = $1
		ORDER BY fecha DESC
	`

	rows, err := r.db.Query(query, idCuenta)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	transacciones := make([]*entities.Transaccion, 0)
	for rows.Next() {
		transaccion := &entities.Transaccion{}
		err := rows.Scan(
			&transaccion.IDTransaccion,
			&transaccion.TipoTransaccion,
			&transaccion.Monto,
			&transaccion.Fecha,
			&transaccion.CuentaOrigen,
			&transaccion.CuentaDestino,
			&transaccion.Descripcion,
		)
		if err != nil {
			return nil, err
		}
		transacciones = append(transacciones, transaccion)
	}

	return transacciones, nil
}
