package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"banco-api/Usuario/application"
)

type GetUsuarioController struct {
	usecase *application.GetUsuarioUseCase
}

func NewGetUsuarioController(usecase *application.GetUsuarioUseCase) *GetUsuarioController {
	return &GetUsuarioController{usecase: usecase}
}

func (ctrl *GetUsuarioController) Handle(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	usuario, err := ctrl.usecase.Execute(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if usuario == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Usuario no encontrado"})
		return
	}

	c.JSON(http.StatusOK, usuario)
}
