package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"banco-api/Transaccion/application"
)

type GetTransactionsByCuentaController struct {
	usecase *application.GetTransactionsByCuentaUseCase
}

func NewGetTransactionsByCuentaController(usecase *application.GetTransactionsByCuentaUseCase) *GetTransactionsByCuentaController {
	return &GetTransactionsByCuentaController{usecase: usecase}
}

func (ctrl *GetTransactionsByCuentaController) Handle(c *gin.Context) {
	idCuenta, err := strconv.Atoi(c.Param("id_cuenta"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID de cuenta inválido"})
		return
	}

	transacciones, err := ctrl.usecase.Execute(idCuenta)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, transacciones)
}
