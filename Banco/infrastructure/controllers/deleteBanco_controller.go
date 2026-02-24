package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"banco-api/Banco/application"
)

type DeleteBancoController struct {
	usecase *application.DeleteBancoUseCase
}

func NewDeleteBancoController(usecase *application.DeleteBancoUseCase) *DeleteBancoController {
	return &DeleteBancoController{usecase: usecase}
}

func (ctrl *DeleteBancoController) Handle(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	if err := ctrl.usecase.Execute(id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"mensaje": "Banco eliminado exitosamente"})
}
