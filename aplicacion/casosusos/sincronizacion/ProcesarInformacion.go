package casosusos_sincronizacion

import (
	"encoding/json"
	"errors"
	"log"
	"ms-sincronizador-tienda/dominio/constantes"
	dominio_repositorios "ms-sincronizador-tienda/dominio/repositorios/db"
	"net/http"
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

	respuesta, err := PI.Procesar.Ejecutar(payload)
	if err != nil {
		log.Println("Error en la sincronizacion de datos", err)
		return err
	}

	if respuesta.CodigoRespuesta != http.StatusCreated {
		log.Println("Error :", respuesta.Mensaje)
		return errors.New("Fallo en procedimiento de sync")
	}
	return nil
}
