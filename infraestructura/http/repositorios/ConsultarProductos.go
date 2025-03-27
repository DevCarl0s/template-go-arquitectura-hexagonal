package repositorios_infraestruture_http

import (
	"encoding/json"
	"errors"
	"log"
	"ms-sincronizador-tienda/dominio/constantes"
	"ms-sincronizador-tienda/dominio/entidades"
	entidades_sincronizacion_productos "ms-sincronizador-tienda/dominio/entidades/sincronizacion/productos"
	dominio_repositorios_http "ms-sincronizador-tienda/dominio/repositorios/http"
	"net/http"
	"strconv"
)

type ConsultarProductos struct {
	Client dominio_repositorios_http.IClienteHttp
}

func (CP *ConsultarProductos) Consultar(peticion *entidades.HttpRequest) (*entidades_sincronizacion_productos.RespuestaProductos, error) {
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
