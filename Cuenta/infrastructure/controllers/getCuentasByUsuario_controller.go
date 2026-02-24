package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"banco-api/Cuenta/application"
)

type GetCuentasByUsuarioController struct {
	usecase *application.GetCuentasByUsuarioUseCase
}

func NewGetCuentasByUsuarioController(usecase *application.GetCuentasByUsuarioUseCase) *GetCuentasByUsuarioController {
	return &GetCuentasByUsuarioController{usecase: usecase}
}

func (ctrl *GetCuentasByUsuarioController) Handle(c *gin.Context) {
	idUsuario, err := strconv.Atoi(c.Param("id_usuario"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID de usuario inválido"})
		return
	}

	cuentas, err := ctrl.usecase.Execute(idUsuario)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, cuentas)
}
