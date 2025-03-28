package casosusos_sincronizacion_consecutivos

import (
	"log"
	"ms-sincronizador-tienda/dominio/constantes"
	"ms-sincronizador-tienda/dominio/entidades"
	entidades_sincronizacion_consecutivos "ms-sincronizador-tienda/dominio/entidades/sincronizacion/consecutivos"
	dominio_repositorios_http "ms-sincronizador-tienda/dominio/repositorios/http"
)

type ConsultarConsecutivos struct {
	Cliente dominio_repositorios_http.IConsultarConsecutivos
}

func (CC *ConsultarConsecutivos) Ejecutar(peticion *entidades.HttpRequest) (*entidades_sincronizacion_consecutivos.RespuestaConsecutivos, error) {
	log.Println(constantes.Green + "[ConsultarConsecutivos] Ejecutar" + constantes.Reset)
	respuesta, err := CC.Cliente.Consultar(peticion)

	if err != nil {
		log.Println(constantes.Red + "Host " + peticion.Url + "sin respuesta")
		log.Println("Error: " + err.Error() + constantes.Reset)
		return nil, err
	}

	return respuesta, nil
}
