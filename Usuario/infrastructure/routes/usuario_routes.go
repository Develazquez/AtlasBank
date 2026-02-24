package routes

import (
	"database/sql"

	"github.com/gin-gonic/gin"

	"banco-api/Usuario/application"
	"banco-api/Usuario/infrastructure/controllers"
	repo "banco-api/Usuario/infrastructure/repository"
)

func SetupUsuarioRoutes(router *gin.Engine, db *sql.DB) {
	usuarioRepo := repo.NewUsuarioRepositoryPostgres(db)


	createUseCase := application.NewCreateUsuarioUseCase(usuarioRepo)
	getUseCase := application.NewGetUsuarioUseCase(usuarioRepo)
	getAllUseCase := application.NewGetAllUsuariosUseCase(usuarioRepo)
	updateUseCase := application.NewUpdateUsuarioUseCase(usuarioRepo)
	deleteUseCase := application.NewDeleteUsuarioUseCase(usuarioRepo)

	createCtrl := controllers.NewCreateUsuarioController(createUseCase)
	getCtrl := controllers.NewGetUsuarioController(getUseCase)
	getAllCtrl := controllers.NewGetAllUsuariosController(getAllUseCase)
	updateCtrl := controllers.NewUpdateUsuarioController(updateUseCase)
	deleteCtrl := controllers.NewDeleteUsuarioController(deleteUseCase)
	loginCtrl := controllers.LoginUsuarioController

	usuarios := router.Group("/atlasApp/usuarios")
	{
		usuarios.POST("", createCtrl.Handle)
		usuarios.GET("", getAllCtrl.Handle)
		usuarios.GET("/:id", getCtrl.Handle)
		usuarios.PUT("/:id", updateCtrl.Handle)
		usuarios.DELETE("/:id", deleteCtrl.Handle)
		usuarios.POST("/login", loginCtrl)
	}
}
