// Usuario/domain/dto/usuario_list_dto.go
package dto

type UsuarioListDTO struct {
    ID              string  `json:"id"`
    CuentaID        *string `json:"cuenta_id"`
    Nombre          string  `json:"nombre"`
    ApellidoPaterno string  `json:"apellido_paterno"`
    Email           string  `json:"email"`
}