package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"banco-api/Banco/application"
)

type GetBancoController struct {
	usecase *application.GetBancoUseCase
}

func NewGetBancoController(usecase *application.GetBancoUseCase) *GetBancoController {
	return &GetBancoController{usecase: usecase}
}

func (ctrl *GetBancoController) Handle(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	banco, err := ctrl.usecase.Execute(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if banco == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Banco no encontrado"})
		return
	}

	c.JSON(http.StatusOK, banco)
}
