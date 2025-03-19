package dominio_repositorios

import entidades_sincronizacion "ms-sincronizador-tienda/dominio/entidades/sincronizacion"

type IProcesarInformacion interface {
	Ejecutar(data []byte) (*entidades_sincronizacion.RespuestaSincronizacion, error)
}
