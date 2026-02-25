package infrastructure

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"banco-api/Transaccion/infrastructure/routes"
)

type TransaccionDependencies struct {
	db *gorm.DB
}

func NewTransaccionDependencies(db *gorm.DB) *TransaccionDependencies {
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
