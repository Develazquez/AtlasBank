package routes

import (
	"database/sql"

	"github.com/gin-gonic/gin"

	"banco-api/Banco/application"
	"banco-api/Banco/infrastructure/controllers"
	repo "banco-api/Banco/infrastructure/repository"
)

func SetupBancoRoutes(router *gin.Engine, db *sql.DB) {
	bancoRepo := repo.NewBancoRepositoryPostgres(db)

	createUseCase := application.NewCreateBancoUseCase(bancoRepo)
	getUseCase := application.NewGetBancoUseCase(bancoRepo)
	getAllUseCase := application.NewGetAllBancosUseCase(bancoRepo)
	updateUseCase := application.NewUpdateBancoUseCase(bancoRepo)
	deleteUseCase := application.NewDeleteBancoUseCase(bancoRepo)


	createCtrl := controllers.NewCreateBancoController(createUseCase)
	getCtrl := controllers.NewGetBancoController(getUseCase)
	getAllCtrl := controllers.NewGetAllBancosController(getAllUseCase)
	updateCtrl := controllers.NewUpdateBancoController(updateUseCase)
	deleteCtrl := controllers.NewDeleteBancoController(deleteUseCase)


	bancos := router.Group("/atlasApp/bancos")
	{
		bancos.POST("", createCtrl.Handle)      
		bancos.GET("", getAllCtrl.Handle)        
		bancos.GET("/:id", getCtrl.Handle)      
		bancos.PUT("/:id", updateCtrl.Handle)    
		bancos.DELETE("/:id", deleteCtrl.Handle)
	}
}
