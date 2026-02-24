package infrastructure

import (
	"database/sql"

	"github.com/gin-gonic/gin"

	"banco-api/Banco/infrastructure/routes"
)

type BancoDependencies struct {
	db *sql.DB
}

func NewBancoDependencies(db *sql.DB) *BancoDependencies {
	return &BancoDependencies{
		db: db,
	}
}

func (bd *BancoDependencies) Setup(router *gin.Engine) {
	routes.SetupBancoRoutes(router, bd.db)
}

func (bd *BancoDependencies) Shutdown() error {
	return nil
}
