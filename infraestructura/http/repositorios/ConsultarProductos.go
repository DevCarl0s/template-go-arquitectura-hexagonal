package repositorios_infraestruture_http

import (
	"encoding/json"
	"errors"
	"log"
	comunes_http_clientes "ms-sincronizador-tienda/comunes/dominio/adaptadores/clientes/http"
	comunes_entidades "ms-sincronizador-tienda/comunes/dominio/entidades"
	"ms-sincronizador-tienda/dominio/constantes"
	entidades_sincronizacion_productos "ms-sincronizador-tienda/dominio/entidades/sincronizacion/productos"
	"net/http"
	"strconv"
)

type ConsultarProductos struct {
	Client comunes_http_clientes.IClienteHttp
}

func (CP *ConsultarProductos) Consultar(peticion *comunes_entidades.HttpRequest) (*entidades_sincronizacion_productos.RespuestaProductos, error) {
	respuesta, err := CP.Client.Enviar(peticion.Metodo, peticion.Url, peticion)
	if err != nil {
		log.Println(constantes.Red + err.Error() + constantes.Reset)
		return nil, err
	}

	if respuesta.StatusCode != http.StatusOK {
		mensaje := "Error al consultar productos: " + strconv.Itoa(respuesta.StatusCode)
		return nil, errors.New(mensaje)
	}

	respuestaProductos := &entidades_sincronizacion_productos.RespuestaProductos{}
	err = json.Unmarshal(respuesta.Body, respuestaProductos)

	if err != nil {
		log.Println(constantes.Red + err.Error() + constantes.Reset)
		return nil, err
	}

	return respuestaProductos, nil
}
