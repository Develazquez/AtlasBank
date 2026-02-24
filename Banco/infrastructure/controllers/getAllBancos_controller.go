package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"banco-api/Banco/application"
)

type GetAllBancosController struct {
	usecase *application.GetAllBancosUseCase
}

func NewGetAllBancosController(usecase *application.GetAllBancosUseCase) *GetAllBancosController {
	return &GetAllBancosController{usecase: usecase}
}

func (ctrl *GetAllBancosController) Handle(c *gin.Context) {
	bancos, err := ctrl.usecase.Execute()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, bancos)
}
