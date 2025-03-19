package comunes_http_clientes

import comunes_entidades "ms-sincronizador-tienda/comunes/dominio/entidades"

type IClienteHttp interface {
	Enviar(metodo string, url string, mensaje *comunes_entidades.HttpRequest) (*comunes_entidades.HttpResponse, error)
}
