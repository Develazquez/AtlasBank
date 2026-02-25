package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"banco-api/Cuenta/application"
	"banco-api/Cuenta/domain/entities"
)

type CreateCuentaController struct {
	usecase *application.CreateCuentaUseCase
}

func NewCreateCuentaController(usecase *application.CreateCuentaUseCase) *CreateCuentaController {
	return &CreateCuentaController{usecase: usecase}
}

func (ctrl *CreateCuentaController) Handle(c *gin.Context) {
	var input struct {
		UsuarioID          string  `json:"usuario_id" binding:"required"`
		BancoID            string  `json:"banco_id" binding:"required"`
		NumeroCuenta       string  `json:"numero_cuenta" binding:"required,min=5"`
		NumeroTarjeta      string  `json:"numero_tarjeta"`
		CLABE              string  `json:"clabe"`
		IBAN               string  `json:"iban"`
		Tipo               string  `json:"tipo" binding:"required"`
		TipoTarjeta        string  `json:"tipo_tarjeta"`
		NombreTarjeta      string  `json:"nombre_tarjeta"`
		FechaExpiracion    string  `json:"fecha_expiracion"`
		Moneda             string  `json:"moneda"`
		Saldo              float64 `json:"saldo"`
		SaldoBloqueado     float64 `json:"saldo_bloqueado"`
		LimiteRetiroDiario float64 `json:"limite_retiro_diario"`
		Estado             string  `json:"estado"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos", "detalles": err.Error()})
		return
	}

	usuarioID, err := uuid.Parse(input.UsuarioID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "usuario_id inválido"})
		return
	}

	bancoID, err := uuid.Parse(input.BancoID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "banco_id inválido"})
		return
	}

	tipoTarjeta := input.TipoTarjeta
	if tipoTarjeta == "" {
		tipoTarjeta = "CLASSIC"
	}

	moneda := input.Moneda
	if moneda == "" {
		moneda = "MXN"
	}

	tipo := input.Tipo
	if tipo == "" {
		tipo = "AHORRO"
	}

	estado := input.Estado
	if estado == "" {
		estado = "ACTIVA"
	}

	limiteRetiroDiario := input.LimiteRetiroDiario
	if limiteRetiroDiario == 0 {
		limiteRetiroDiario = 10000.00
	}

	cuenta := &entities.Cuenta{
		ID:                 uuid.New(),
		UsuarioID:          usuarioID,
		BancoID:            bancoID,
		NumeroCuenta:       input.NumeroCuenta,
		NumeroTarjeta:      input.NumeroTarjeta,
		CLABE:              input.CLABE,
		IBAN:               input.IBAN,
		Tipo:               entities.TipoCuenta(tipo),
		TipoTarjeta:        entities.TipoTarjeta(tipoTarjeta),
		NombreTarjeta:      input.NombreTarjeta,
		FechaExpiracion:    input.FechaExpiracion,
		Moneda:             entities.Moneda(moneda),
		Saldo:              input.Saldo,
		SaldoBloqueado:     input.SaldoBloqueado,
		LimiteRetiroDiario: limiteRetiroDiario,
		Estado:             entities.EstadoCuenta(estado),
		FechaApertura:      time.Now(),
	}

	id, err := ctrl.usecase.Execute(cuenta)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"mensaje": "Cuenta creada exitosamente",
		"id":      id,
	})
}
