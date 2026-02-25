package infrastructure

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"banco-api/Cuenta/infrastructure/routes"
)

type CuentaDependencies struct {
	db *gorm.DB
}

func NewCuentaDependencies(db *gorm.DB) *CuentaDependencies {
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
