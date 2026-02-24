package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"banco-api/Transaccion/application"
)

type DeleteTransaccionController struct {
	usecase *application.DeleteTransaccionUseCase
}

func NewDeleteTransaccionController(usecase *application.DeleteTransaccionUseCase) *DeleteTransaccionController {
	return &DeleteTransaccionController{usecase: usecase}
}

func (ctrl *DeleteTransaccionController) Handle(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	if err := ctrl.usecase.Execute(id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"mensaje": "Transacción eliminada exitosamente"})
}
