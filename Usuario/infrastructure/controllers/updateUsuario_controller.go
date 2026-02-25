package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"banco-api/Usuario/application"
	"banco-api/Usuario/domain/entities"
)

type UpdateUsuarioController struct {
	usecase *application.UpdateUsuarioUseCase
}

func NewUpdateUsuarioController(usecase *application.UpdateUsuarioUseCase) *UpdateUsuarioController {
	return &UpdateUsuarioController{usecase: usecase}
}

type UpdateUsuarioInput struct {
	Nombre          string `json:"nombre"`
	ApellidoPaterno string `json:"apellido_paterno"`
	ApellidoMaterno string `json:"apellido_materno"`
	Email           string `json:"email"`
	Telefono        string `json:"telefono"`
	FechaNacimiento string `json:"fecha_nacimiento"`
}

func (ctrl *UpdateUsuarioController) Handle(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var input UpdateUsuarioInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos", "detalles": err.Error()})
		return
	}

	fecha, err := time.Parse("2006-01-02", input.FechaNacimiento)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":    "Formato de fecha inválido",
			"detalles": "Se esperaba YYYY-MM-DD",
		})
		return
	}

	usuario := entities.Usuario{
		ID:              id,
		Nombre:          input.Nombre,
		ApellidoPaterno: input.ApellidoPaterno,
		ApellidoMaterno: input.ApellidoMaterno,
		Email:           input.Email,
		Telefono:        input.Telefono,
		FechaNacimiento: fecha,
	}

	if err := ctrl.usecase.Execute(&usuario); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"mensaje": "Usuario actualizado exitosamente"})
}
