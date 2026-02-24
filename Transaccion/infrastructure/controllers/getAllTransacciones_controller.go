package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"banco-api/Transaccion/application"
)

type GetAllTransaccionesController struct {
	usecase *application.GetAllTransaccionesUseCase
}

func NewGetAllTransaccionesController(usecase *application.GetAllTransaccionesUseCase) *GetAllTransaccionesController {
	return &GetAllTransaccionesController{usecase: usecase}
}

func (ctrl *GetAllTransaccionesController) Handle(c *gin.Context) {
	transacciones, err := ctrl.usecase.Execute()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, transacciones)
}
