package infrastructure

import (
	"database/sql"

	"github.com/gin-gonic/gin"

	"banco-api/Usuario/infrastructure/routes"
)

type UsuarioDependencies struct {
	db *sql.DB
}

func NewUsuarioDependencies(db *sql.DB) *UsuarioDependencies {
	return &UsuarioDependencies{
		db: db,
	}
}

func (ud *UsuarioDependencies) Setup(router *gin.Engine) {
	routes.SetupUsuarioRoutes(router, ud.db)
}

func (ud *UsuarioDependencies) Shutdown() error {
	return nil
}
