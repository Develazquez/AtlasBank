package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"banco-api/Banco/application"
	"banco-api/Banco/domain/entities"
)

type CreateBancoController struct {
	usecase *application.CreateBancoUseCase
}

func NewCreateBancoController(usecase *application.CreateBancoUseCase) *CreateBancoController {
	return &CreateBancoController{usecase: usecase}
}

func (ctrl *CreateBancoController) Handle(c *gin.Context) {
	var banco entities.Banco
	if err := c.ShouldBindJSON(&banco); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos", "detalles": err.Error()})
		return
	}

	createdBanco, err := ctrl.usecase.Execute(&banco)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"mensaje": "Banco creado exitosamente",
		"data":    createdBanco,
	})
}
