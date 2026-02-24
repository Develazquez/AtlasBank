package controllers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"

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
	Apellido        string `json:"apellido"`
	Email           string `json:"email"`
	Telefono        string `json:"telefono"`
	FechaNacimiento string `json:"fecha_nacimiento"` 
}

func (ctrl *UpdateUsuarioController) Handle(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
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
		IDUsuario:       id,
		Nombre:          input.Nombre,
		Apellido:        input.Apellido,
		Email:           input.Email,
		Telefono:        input.Telefono,
		FechaNacimiento: &fecha, 
	}

	if err := ctrl.usecase.Execute(&usuario); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"mensaje": "Usuario actualizado exitosamente"})
}