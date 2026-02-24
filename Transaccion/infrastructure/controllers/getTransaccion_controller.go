package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"banco-api/Transaccion/application"
)

type GetTransaccionController struct {
	usecase *application.GetTransaccionUseCase
}

func NewGetTransaccionController(usecase *application.GetTransaccionUseCase) *GetTransaccionController {
	return &GetTransaccionController{usecase: usecase}
}

func (ctrl *GetTransaccionController) Handle(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	transaccion, err := ctrl.usecase.Execute(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if transaccion == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Transacción no encontrada"})
		return
	}

	c.JSON(http.StatusOK, transaccion)
}
