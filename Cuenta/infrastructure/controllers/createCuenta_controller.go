package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"banco-api/Cuenta/application"
	"banco-api/Cuenta/domain/entities"
)

type CreateCuentaController struct {
	usecase *application.CreateCuentaUseCase
}

func NewCreateCuentaController(usecase *application.CreateCuentaUseCase) *CreateCuentaController {
	return &CreateCuentaController{usecase: usecase}
}

func (ctrl *CreateCuentaController) Handle(c *gin.Context) {
	var cuenta entities.Cuenta
	if err := c.ShouldBindJSON(&cuenta); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos", "detalles": err.Error()})
		return
	}

	id, err := ctrl.usecase.Execute(&cuenta)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"mensaje": "Cuenta creada exitosamente",
		"id":      id,
	})
}
