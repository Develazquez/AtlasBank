package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"banco-api/Usuario/application"
	"banco-api/Usuario/domain/entities"
)

type UpdateUsuarioController struct {
	usecase *application.UpdateUsuarioUseCase
}

func NewUpdateUsuarioController(usecase *application.UpdateUsuarioUseCase) *UpdateUsuarioController {
	return &UpdateUsuarioController{usecase: usecase}
}

func (ctrl *UpdateUsuarioController) Handle(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var usuario entities.Usuario
	if err := c.ShouldBindJSON(&usuario); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos", "detalles": err.Error()})
		return
	}

	usuario.IDUsuario = id
	if err := ctrl.usecase.Execute(&usuario); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"mensaje": "Usuario actualizado exitosamente"})
}
