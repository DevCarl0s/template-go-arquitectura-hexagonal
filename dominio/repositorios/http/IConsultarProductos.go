package dominio_repositorios_http

import (
	"ms-sincronizador-tienda/dominio/entidades"
	entidades_sincronizacion_productos "ms-sincronizador-tienda/dominio/entidades/sincronizacion/productos"
)

type IConsultarProductos interface {
	Consultar(peticion *entidades.HttpRequest) (*entidades_sincronizacion_productos.RespuestaProductos, error)
}
