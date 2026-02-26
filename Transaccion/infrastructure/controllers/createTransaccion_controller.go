package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"banco-api/Transaccion/application"
	"banco-api/Transaccion/domain/entities"
	"banco-api/Transaccion/infrastructure/websocket"
)

type CreateTransaccionController struct {
	usecase *application.CreateTransaccionUseCase
	db      *sql.DB
}

func NewCreateTransaccionController(usecase *application.CreateTransaccionUseCase, db *sql.DB) *CreateTransaccionController {
	return &CreateTransaccionController{usecase: usecase, db: db}
}

func (ctrl *CreateTransaccionController) Handle(c *gin.Context) {
	var input struct {
		CuentaOrigenID  *string `json:"cuenta_origen_id"`
		CuentaDestinoID *string `json:"cuenta_destino_id"`
		Monto           float64 `json:"monto" binding:"required,gt=0"`
		Concepto        string  `json:"concepto"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos", "detalles": err.Error()})
		return
	}

	var cuentaOrigenID, cuentaDestinoID *uuid.UUID
	if input.CuentaOrigenID != nil {
		id, err := uuid.Parse(*input.CuentaOrigenID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "cuenta_origen_id inválido"})
			return
		}
		cuentaOrigenID = &id
	}

	if input.CuentaDestinoID != nil {
		id, err := uuid.Parse(*input.CuentaDestinoID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "cuenta_destino_id inválido"})
			return
		}
		cuentaDestinoID = &id
	}

	transaccion := &entities.Transaccion{
		ID:              uuid.New(),
		CuentaOrigenID:  cuentaOrigenID,
		CuentaDestinoID: cuentaDestinoID,
		Tipo:            entities.TRANSFERENCIA_INTERNA,
		Estado:          entities.COMPLETADA,
		Monto:           input.Monto,
		Moneda:          "MXN",
		Concepto:        input.Concepto,
		Referencia:      uuid.New().String(),
		Canal:           entities.APP,
		CreatedAt:       time.Now(),
	}

	transaccionJSON, _ := json.MarshalIndent(transaccion, "", "  ")
	fmt.Printf("Transacción creada: %s\n", string(transaccionJSON))

	id, err := ctrl.usecase.Execute(transaccion)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if transaccion.CuentaOrigenID != nil && transaccion.CuentaDestinoID != nil {
		fmt.Printf("estoy aqui")
		go websocket.NotifyTransfer(true)
	}

	c.JSON(http.StatusCreated, gin.H{
		"mensaje": "Transacción realizada exitosamente",
		"id":      id,
	})
}