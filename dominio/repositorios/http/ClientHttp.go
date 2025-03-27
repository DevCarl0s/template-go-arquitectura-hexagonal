package dominio_repositorios_http

import "ms-sincronizador-tienda/dominio/entidades"

type IClienteHttp interface {
	Enviar(metodo string, url string, mensaje *entidades.HttpRequest) (*entidades.HttpResponse, error)
}
