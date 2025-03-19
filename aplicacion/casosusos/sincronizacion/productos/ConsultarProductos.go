package casosusos_sincronizacion_productos

import (
	"log"
	comunes_entidades "ms-sincronizador-tienda/comunes/dominio/entidades"
	"ms-sincronizador-tienda/dominio/constantes"
	entidades_sincronizacion_productos "ms-sincronizador-tienda/dominio/entidades/sincronizacion/productos"
	dominio_repositorios_http "ms-sincronizador-tienda/dominio/repositorios/http"
)

type ConsultarProductos struct {
	Cliente dominio_repositorios_http.IConsultarProductos
}

func (CP *ConsultarProductos) Ejecutar(peticion *comunes_entidades.HttpRequest) (*entidades_sincronizacion_productos.RespuestaProductos, error) {
	log.Println(constantes.Green + "[ConsultarProductos] Ejecutar" + constantes.Reset)
	// respuesta, _ := CP.Cliente.Consultar(peticion)

	// if err != nil {
	// 	log.Println(constantes.Red + "Host " + peticion.Url + "sin respuesta")
	// 	log.Println("Error: " + err.Error() + constantes.Reset)
	// 	return nil, err
	// }

	respuesta := &entidades_sincronizacion_productos.RespuestaProductos{
		Id: 1,
	}
	return respuesta, nil
}
