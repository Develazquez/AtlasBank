package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"banco-api/Cuenta/application"
)

type DeleteCuentaController struct {
	usecase *application.DeleteCuentaUseCase
}

func NewDeleteCuentaController(usecase *application.DeleteCuentaUseCase) *DeleteCuentaController {
	return &DeleteCuentaController{usecase: usecase}
}

func (ctrl *DeleteCuentaController) Handle(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	if err := ctrl.usecase.Execute(id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"mensaje": "Cuenta eliminada exitosamente"})
}
