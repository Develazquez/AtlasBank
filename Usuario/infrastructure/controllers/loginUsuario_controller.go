package controllers

import (
	"banco-api/Usuario/application"
	"net/http"

	"github.com/gin-gonic/gin"
)

type LoginUsuarioController struct {
	usecase *application.LoginUsuarioUseCase
}

func NewLoginUsuarioController(usecase *application.LoginUsuarioUseCase) *LoginUsuarioController {
	return &LoginUsuarioController{usecase: usecase}
}

func (ctrl *LoginUsuarioController) Handle(c *gin.Context) {
	var req struct {
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos", "detalles": err.Error()})
		return
	}

	usuario, err := ctrl.usecase.Execute(req.Email, req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Credenciales incorrectas"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"mensaje": "Login exitoso",
		"usuario": usuario,
	})
}
