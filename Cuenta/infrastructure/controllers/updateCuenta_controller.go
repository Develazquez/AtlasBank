package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"banco-api/Cuenta/application"
	"banco-api/Cuenta/domain/entities"
)

type UpdateCuentaController struct {
	usecase *application.UpdateCuentaUseCase
}

func NewUpdateCuentaController(usecase *application.UpdateCuentaUseCase) *UpdateCuentaController {
	return &UpdateCuentaController{usecase: usecase}
}

func (ctrl *UpdateCuentaController) Handle(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var cuenta entities.Cuenta
	if err := c.ShouldBindJSON(&cuenta); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos", "detalles": err.Error()})
		return
	}

	cuenta.ID = id
	if err := ctrl.usecase.Execute(&cuenta); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"mensaje": "Cuenta actualizada exitosamente"})
}
