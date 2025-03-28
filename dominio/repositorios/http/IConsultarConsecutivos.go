package dominio_repositorios_http

import (
	"ms-sincronizador-tienda/dominio/entidades"
	entidades_sincronizacion_consecutivos "ms-sincronizador-tienda/dominio/entidades/sincronizacion/consecutivos"
)

type IConsultarConsecutivos interface {
	Consultar(peticion *entidades.HttpRequest) (*entidades_sincronizacion_consecutivos.RespuestaConsecutivos, error)
}
