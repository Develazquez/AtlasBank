package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"banco-api/Banco/application"
	"banco-api/Banco/domain/entities"
)

type UpdateBancoController struct {
	usecase *application.UpdateBancoUseCase
}

func NewUpdateBancoController(usecase *application.UpdateBancoUseCase) *UpdateBancoController {
	return &UpdateBancoController{usecase: usecase}
}

func (ctrl *UpdateBancoController) Handle(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var banco entities.Banco
	if err := c.ShouldBindJSON(&banco); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos", "detalles": err.Error()})
		return
	}

	banco.ID = id
	if err := ctrl.usecase.Execute(&banco); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"mensaje": "Banco actualizado exitosamente"})
}
