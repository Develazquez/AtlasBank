package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"banco-api/Transaccion/application"
	"banco-api/Transaccion/infrastructure/controllers"
	transaccionRepo "banco-api/Transaccion/infrastructure/repository"
	cuentaRepo "banco-api/Cuenta/infrastructure/repository" 
)

func SetupTransaccionRoutes(router *gin.Engine, db *gorm.DB) {
	transRepo := transaccionRepo.NewTransaccionRepositoryPostgres(db)
	cuentRepo := cuentaRepo.NewCuentaRepositoryPostgres(db) 

	createUseCase := application.NewCreateTransaccionUseCase(transRepo, cuentRepo, db) 
	getUseCase := application.NewGetTransaccionUseCase(transRepo)
	getAllUseCase := application.NewGetAllTransaccionesUseCase(transRepo)
	getTransactionsByCuentaUseCase := application.NewGetTransactionsByCuentaUseCase(transRepo)
	deleteUseCase := application.NewDeleteTransaccionUseCase(transRepo)

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