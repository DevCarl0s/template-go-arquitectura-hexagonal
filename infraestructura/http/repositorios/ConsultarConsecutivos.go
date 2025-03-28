package repositorios_infraestruture_http

import (
	"encoding/json"
	"errors"
	"log"
	"ms-sincronizador-tienda/dominio/constantes"
	"ms-sincronizador-tienda/dominio/entidades"
	entidades_sincronizacion_consecutivos "ms-sincronizador-tienda/dominio/entidades/sincronizacion/consecutivos"
	dominio_repositorios_http "ms-sincronizador-tienda/dominio/repositorios/http"
	"net/http"
	"strconv"
)

type ConsultarConsecutivos struct {
	Client dominio_repositorios_http.IClienteHttp
}

func (CP *ConsultarConsecutivos) Consultar(peticion *entidades.HttpRequest) (*entidades_sincronizacion_consecutivos.RespuestaConsecutivos, error) {
	respuesta, err := CP.Client.Enviar(peticion.Metodo, peticion.Url, peticion)
	if err != nil {
		log.Println(constantes.Red + err.Error() + constantes.Reset)
		return nil, err
	}

	if respuesta.StatusCode != http.StatusOK {
		mensaje := "Error al consultar productos: " + strconv.Itoa(respuesta.StatusCode)
		return nil, errors.New(mensaje)
	}

	respuestaProductos := &entidades_sincronizacion_consecutivos.RespuestaConsecutivos{}
	err = json.Unmarshal(respuesta.Body, respuestaProductos)

	if err != nil {
		log.Println(constantes.Red + err.Error() + constantes.Reset)
		return nil, err
	}

	return respuestaProductos, nil
}
