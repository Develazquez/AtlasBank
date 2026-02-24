package controllers

import (
	"banco-api/Usuario/application"
	"net/http"

	"github.com/gin-gonic/gin"
)

func LoginUsuarioController(c *gin.Context) {
	var req struct {
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
		return
	}

	usuario, err := application.LoginUsuarioUsecase(req.Email, req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Credenciales incorrectas"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"usuario": usuario})
}
