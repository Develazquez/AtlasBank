package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"banco-api/Cuenta/application"
)

type GetCuentaController struct {
	usecase *application.GetCuentaUseCase
}

func NewGetCuentaController(usecase *application.GetCuentaUseCase) *GetCuentaController {
	return &GetCuentaController{usecase: usecase}
}

func (ctrl *GetCuentaController) Handle(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	cuenta, err := ctrl.usecase.Execute(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if cuenta == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Cuenta no encontrada"})
		return
	}

	c.JSON(http.StatusOK, cuenta)
}
