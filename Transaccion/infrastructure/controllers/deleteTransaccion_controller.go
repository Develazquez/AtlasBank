package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"banco-api/Transaccion/application"
)

type DeleteTransaccionController struct {
	usecase *application.DeleteTransaccionUseCase
}

func NewDeleteTransaccionController(usecase *application.DeleteTransaccionUseCase) *DeleteTransaccionController {
	return &DeleteTransaccionController{usecase: usecase}
}

func (ctrl *DeleteTransaccionController) Handle(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
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
