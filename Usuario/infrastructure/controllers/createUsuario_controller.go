package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"banco-api/Usuario/application"
	"banco-api/Usuario/domain/entities"
	"time"
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
        Apellido        string `json:"apellido"`
        Email           string `json:"email"`
        Telefono        string `json:"telefono"`
        FechaNacimiento string `json:"fecha_nacimiento"` 
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

   
    usuario := entities.Usuario{
        Nombre:          input.Nombre,
        Apellido:        input.Apellido,
        Email:           input.Email,
        FechaNacimiento: &fecha, 
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
