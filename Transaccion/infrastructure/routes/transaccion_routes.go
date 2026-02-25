package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"banco-api/Transaccion/application"
	"banco-api/Transaccion/infrastructure/controllers"
	repo "banco-api/Transaccion/infrastructure/repository"
)

func SetupTransaccionRoutes(router *gin.Engine, db *gorm.DB) {
	transaccionRepo := repo.NewTransaccionRepositoryPostgres(db)

	createUseCase := application.NewCreateTransaccionUseCase(transaccionRepo, db)
	getUseCase := application.NewGetTransaccionUseCase(transaccionRepo)
	getAllUseCase := application.NewGetAllTransaccionesUseCase(transaccionRepo)
	getTransactionsByCuentaUseCase := application.NewGetTransactionsByCuentaUseCase(transaccionRepo)
	deleteUseCase := application.NewDeleteTransaccionUseCase(transaccionRepo)

	createSQLDB, err := db.DB()
	if err != nil {
		return
	}

	createCtrl := controllers.NewCreateTransaccionController(createUseCase, createSQLDB)
	getCtrl := controllers.NewGetTransaccionController(getUseCase)
	getAllCtrl := controllers.NewGetAllTransaccionesController(getAllUseCase)
	getTransactionsByCuentaCtrl := controllers.NewGetTransactionsByCuentaController(getTransactionsByCuentaUseCase)
	deleteCtrl := controllers.NewDeleteTransaccionController(deleteUseCase)

	transacciones := router.Group("/atlasApp/transacciones")
	{
		transacciones.POST("", createCtrl.Handle)
		transacciones.GET("", getAllCtrl.Handle)
		transacciones.GET("/:id", getCtrl.Handle)
		transacciones.GET("/cuenta/:id_cuenta", getTransactionsByCuentaCtrl.Handle)
		transacciones.DELETE("/:id", deleteCtrl.Handle)
	}
}
