package repositorios_infraestruture

import (
	"log"
	comunes_db_clientes "ms-sincronizador-tienda/comunes/dominio/adaptadores/clientes/db"
	entidades_sincronizacion "ms-sincronizador-tienda/dominio/entidades/sincronizacion"
	"strconv"
)

type ProcesarInformacion struct {
	Cliente comunes_db_clientes.IClienteDB
}

func (PI *ProcesarInformacion) Ejecutar(data []byte) (*entidades_sincronizacion.RespuestaSincronizacion, error) {

	respuesta, err := PI.Cliente.Select(`CALL sincronizacion.prc_procesar_sincronizacion_tienda($1, '{}')`, []any{string(data)})

	if err != nil {
		return nil, err
	}

	log.Println("Respuesta prc_procesar_sincronizacion: ", respuesta)
	parametro := &entidades_sincronizacion.RespuestaSincronizacion{}
	for _, valor := range respuesta {
		if datos, ok := valor[0].(map[string]interface{}); ok {
			if codigo, existe := datos["codigo_respuesta"].(string); existe {
				codigoError, _ := strconv.Atoi(codigo)
				parametro.CodigoRespuesta = codigoError
			} else {
				parametro.CodigoRespuesta = int(datos["codigo_respuesta"].(float64))
			}
			if mensaje, existe := datos["estado"].(string); existe {
				parametro.Mensaje = mensaje
			}
		}
	}

	log.Println(parametro)
	return parametro, nil
}
