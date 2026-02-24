package routes

import (
	"database/sql"

	"github.com/gin-gonic/gin"

	"banco-api/Cuenta/application"
	"banco-api/Cuenta/infrastructure/controllers"
	repo "banco-api/Cuenta/infrastructure/repository"
)

func SetupCuentaRoutes(router *gin.Engine, db *sql.DB) {
	cuentaRepo := repo.NewCuentaRepositoryPostgres(db)

	createUseCase := application.NewCreateCuentaUseCase(cuentaRepo)
	getUseCase := application.NewGetCuentaUseCase(cuentaRepo)
	getAllUseCase := application.NewGetAllCuentasUseCase(cuentaRepo)
	getCuentasByUsuarioUseCase := application.NewGetCuentasByUsuarioUseCase(cuentaRepo)
	updateUseCase := application.NewUpdateCuentaUseCase(cuentaRepo)
	deleteUseCase := application.NewDeleteCuentaUseCase(cuentaRepo)

	createCtrl := controllers.NewCreateCuentaController(createUseCase)
	getCtrl := controllers.NewGetCuentaController(getUseCase)
	getAllCtrl := controllers.NewGetAllCuentasController(getAllUseCase)
	getCuentasByUsuarioCtrl := controllers.NewGetCuentasByUsuarioController(getCuentasByUsuarioUseCase)
	updateCtrl := controllers.NewUpdateCuentaController(updateUseCase)
	deleteCtrl := controllers.NewDeleteCuentaController(deleteUseCase)

	cuentas := router.Group("/api/v1/cuentas")
	{
		cuentas.POST("", createCtrl.Handle)
		cuentas.GET("", getAllCtrl.Handle)
		cuentas.GET("/:id", getCtrl.Handle)
		cuentas.GET("/usuario/:id_usuario", getCuentasByUsuarioCtrl.Handle)
		cuentas.PUT("/:id", updateCtrl.Handle)
		cuentas.DELETE("/:id", deleteCtrl.Handle)
	}
}
