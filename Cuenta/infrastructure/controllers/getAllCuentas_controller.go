package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"banco-api/Cuenta/application"
)

type GetAllCuentasController struct {
	usecase *application.GetAllCuentasUseCase
}

func NewGetAllCuentasController(usecase *application.GetAllCuentasUseCase) *GetAllCuentasController {
	return &GetAllCuentasController{usecase: usecase}
}

func (ctrl *GetAllCuentasController) Handle(c *gin.Context) {
	cuentas, err := ctrl.usecase.Execute()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, cuentas)
}
