package infrastructure

import (
	"database/sql"

	"github.com/gin-gonic/gin"

	"banco-api/Cuenta/infrastructure/routes"
)

type CuentaDependencies struct {
	db *sql.DB
}

func NewCuentaDependencies(db *sql.DB) *CuentaDependencies {
	return &CuentaDependencies{
		db: db,
	}
}

func (cd *CuentaDependencies) Setup(router *gin.Engine) {
	routes.SetupCuentaRoutes(router, cd.db)
}

func (cd *CuentaDependencies) Shutdown() error {
	return nil
}
