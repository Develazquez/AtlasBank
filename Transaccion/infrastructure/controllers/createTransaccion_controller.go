package controllers

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

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
	var input struct {
		CuentaOrigenID  *string                `json:"cuenta_origen_id"`
		CuentaDestinoID *string                `json:"cuenta_destino_id"`
		Tipo            string                 `json:"tipo" binding:"required"`
		Estado          string                 `json:"estado"`
		Monto           float64                `json:"monto" binding:"required,gt=0"`
		Moneda          string                 `json:"moneda"`
		Comision        float64                `json:"comision"`
		Concepto        string                 `json:"concepto"`
		Descripcion     string                 `json:"descripcion"`
		Referencia      string                 `json:"referencia"`
		IPOrigen        string                 `json:"ip_origen"`
		Canal           string                 `json:"canal"`
		Metadata        map[string]interface{} `json:"metadata"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos", "detalles": err.Error()})
		return
	}

	// Parsear UUIDs opcionales
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

	// Asignar valores por defecto
	estado := input.Estado
	if estado == "" {
		estado = "PENDIENTE"
	}

	moneda := input.Moneda
	if moneda == "" {
		moneda = "MXN"
	}

	canal := input.Canal
	if canal == "" {
		canal = "APP"
	}

	referencia := input.Referencia
	if referencia == "" {
		referencia = uuid.New().String()
	}

	transaccion := &entities.Transaccion{
		ID:              uuid.New(),
		CuentaOrigenID:  cuentaOrigenID,
		CuentaDestinoID: cuentaDestinoID,
		Tipo:            entities.TipoTransaccion(input.Tipo),
		Estado:          entities.EstadoTransaccion(estado),
		Monto:           input.Monto,
		Moneda:          moneda,
		Comision:        input.Comision,
		Concepto:        input.Concepto,
		Descripcion:     input.Descripcion,
		Referencia:      referencia,
		IPOrigen:        input.IPOrigen,
		Canal:           entities.Canal(canal),
		Metadata:        input.Metadata,
		CreatedAt:       time.Now(),
	}

	id, err := ctrl.usecase.Execute(transaccion)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"mensaje": "Transacción realizada exitosamente",
		"id":      id,
	})
}
