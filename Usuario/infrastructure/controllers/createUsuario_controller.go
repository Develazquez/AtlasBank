package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"banco-api/Usuario/application"
	"banco-api/Usuario/domain/entities"
)

type CreateUsuarioController struct {
	usecase *application.CreateUsuarioUseCase
}

func NewCreateUsuarioController(usecase *application.CreateUsuarioUseCase) *CreateUsuarioController {
	return &CreateUsuarioController{usecase: usecase}
}

func (ctrl *CreateUsuarioController) Handle(c *gin.Context) {
	var usuario entities.Usuario
	if err := c.ShouldBindJSON(&usuario); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos", "detalles": err.Error()})
		return
	}

	id, err := ctrl.usecase.Execute(&usuario)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"mensaje": "Usuario creado exitosamente",
		"id":      id,
	})
}
