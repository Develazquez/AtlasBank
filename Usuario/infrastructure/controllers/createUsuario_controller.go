package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"banco-api/Usuario/application"
	"banco-api/Usuario/domain/entities"
)

type CreateUsuarioController struct {
	usecase *application.CreateUsuarioUseCase
}

func NewCreateUsuarioController(usecase *application.CreateUsuarioUseCase) *CreateUsuarioController {
	return &CreateUsuarioController{usecase: usecase}
}

func (ctrl *CreateUsuarioController) Handle(c *gin.Context) {
	var input struct {
		Nombre          string `json:"nombre"`
		ApellidoPaterno string `json:"apellido_paterno"`
		ApellidoMaterno string `json:"apellido_materno"`
		Email           string `json:"email"`
		Telefono        string `json:"telefono"`
		FechaNacimiento string `json:"fecha_nacimiento"`
		BancoID         string `json:"banco_id"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos", "detalles": err.Error()})
		return
	}

	fecha, err := time.Parse("2006-01-02", input.FechaNacimiento)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Formato de fecha inválido. Use YYYY-MM-DD"})
		return
	}

	bancoID, err := uuid.Parse(input.BancoID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "banco_id inválido"})
		return
	}

	usuario := entities.Usuario{
		ID:              uuid.New(),
		BancoID:         bancoID,
		Nombre:          input.Nombre,
		ApellidoPaterno: input.ApellidoPaterno,
		ApellidoMaterno: input.ApellidoMaterno,
		Email:           input.Email,
		Telefono:        input.Telefono,
		FechaNacimiento: fecha,
	}

	id, err := ctrl.usecase.Execute(&usuario)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"mensaje": "Usuario creado exitosamente",
		"id":      id,
	})
}
