package casosusos_sincronizacion

import (
	"encoding/json"
	"log"
	"ms-sincronizador-tienda/dominio/constantes"
	dominio_repositorios "ms-sincronizador-tienda/dominio/repositorios/db"
)

type ProcesarInformacion struct {
	Procesar dominio_repositorios.IProcesarInformacion
}

func (PI *ProcesarInformacion) ObtenerPayload(data any, descripcion string) ([]byte, error) {
	payload := map[string]any{
		descripcion: data,
	}
	jsonData, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}
	return jsonData, nil
}

func (PI *ProcesarInformacion) Ejecutar(data any, descripcion string) error {
	log.Println(constantes.Blue + "Sincronizando " + descripcion + constantes.Reset)
	payload, err := PI.ObtenerPayload(data, descripcion)
	if err != nil {
		log.Println("Error al convertir data a sincronizar " + "[" + descripcion + "]")
		return err
	}
	log.Println("Datos a procesar: ", string(payload))

	// Llamar Procedimiento
	PI.Procesar.Ejecutar(payload)
	return nil
}
