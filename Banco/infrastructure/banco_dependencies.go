package infrastructure

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"banco-api/Banco/infrastructure/routes"
)

type BancoDependencies struct {
	db *gorm.DB
}

func NewBancoDependencies(db *gorm.DB) *BancoDependencies {
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
