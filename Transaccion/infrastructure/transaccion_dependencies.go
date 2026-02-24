package infrastructure

import (
	"database/sql"

	"github.com/gin-gonic/gin"

	"banco-api/Transaccion/infrastructure/routes"
)

type TransaccionDependencies struct {
	db *sql.DB
}

func NewTransaccionDependencies(db *sql.DB) *TransaccionDependencies {
	return &TransaccionDependencies{
		db: db,
	}
}

func (td *TransaccionDependencies) Setup(router *gin.Engine) {
	routes.SetupTransaccionRoutes(router, td.db)
}

func (td *TransaccionDependencies) Shutdown() error {
	return nil
}
