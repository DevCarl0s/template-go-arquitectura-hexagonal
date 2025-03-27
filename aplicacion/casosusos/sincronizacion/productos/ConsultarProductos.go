package casosusos_sincronizacion_productos

import (
	"log"
	"ms-sincronizador-tienda/dominio/constantes"
	"ms-sincronizador-tienda/dominio/entidades"
	entidades_sincronizacion_productos "ms-sincronizador-tienda/dominio/entidades/sincronizacion/productos"
	dominio_repositorios_http "ms-sincronizador-tienda/dominio/repositorios/http"
)

type ConsultarProductos struct {
	Cliente dominio_repositorios_http.IConsultarProductos
}

func (CP *ConsultarProductos) Ejecutar(peticion *entidades.HttpRequest) (*entidades_sincronizacion_productos.RespuestaProductos, error) {
	log.Println(constantes.Green + "[ConsultarProductos] Ejecutar" + constantes.Reset)
	respuesta, err := CP.Cliente.Consultar(peticion)

	if err != nil {
		log.Println(constantes.Red + "Host " + peticion.Url + "sin respuesta")
		log.Println("Error: " + err.Error() + constantes.Reset)
		return nil, err
	}

	return respuesta, nil
}
