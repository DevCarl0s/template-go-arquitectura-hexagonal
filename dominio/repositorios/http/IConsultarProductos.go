package dominio_repositorios_http

import (
	comunes_entidades "ms-sincronizador-tienda/comunes/dominio/entidades"
	entidades_sincronizacion_productos "ms-sincronizador-tienda/dominio/entidades/sincronizacion/productos"
)

type IConsultarProductos interface {
	Consultar(peticion *comunes_entidades.HttpRequest) (*entidades_sincronizacion_productos.RespuestaProductos, error)
}
