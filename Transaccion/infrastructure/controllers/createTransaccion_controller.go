package controllers

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"

	"banco-api/Transaccion/application"
	"banco-api/Transaccion/domain/entities"
)

type CreateTransaccionController struct {
	usecase *application.CreateTransaccionUseCase
	db      *sql.DB
}

func NewCreateTransaccionController(usecase *application.CreateTransaccionUseCase, db *sql.DB) *CreateTransaccionController {
	return &CreateTransaccionController{usecase: usecase, db: db}
}

func (ctrl *CreateTransaccionController) Handle(c *gin.Context) {
	var transaccion entities.Transaccion
	if err := c.ShouldBindJSON(&transaccion); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos", "detalles": err.Error()})
		return
	}

	id, err := ctrl.usecase.Execute(&transaccion)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"mensaje": "Transacción realizada exitosamente",
		"id":      id,
	})
}
