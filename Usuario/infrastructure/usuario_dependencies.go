package infrastructure

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"banco-api/Usuario/infrastructure/routes"
)

type UsuarioDependencies struct {
	db *gorm.DB
}

func NewUsuarioDependencies(db *gorm.DB) *UsuarioDependencies {
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
