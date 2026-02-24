package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"banco-api/Usuario/application"
)

type GetAllUsuariosController struct {
	usecase *application.GetAllUsuariosUseCase
}

func NewGetAllUsuariosController(usecase *application.GetAllUsuariosUseCase) *GetAllUsuariosController {
	return &GetAllUsuariosController{usecase: usecase}
}

func (ctrl *GetAllUsuariosController) Handle(c *gin.Context) {
	usuarios, err := ctrl.usecase.Execute()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, usuarios)
}
