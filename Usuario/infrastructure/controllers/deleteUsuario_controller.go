package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"banco-api/Usuario/application"
)

type DeleteUsuarioController struct {
	usecase *application.DeleteUsuarioUseCase
}

func NewDeleteUsuarioController(usecase *application.DeleteUsuarioUseCase) *DeleteUsuarioController {
	return &DeleteUsuarioController{usecase: usecase}
}

func (ctrl *DeleteUsuarioController) Handle(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	if err := ctrl.usecase.Execute(id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"mensaje": "Usuario eliminado exitosamente"})
}
