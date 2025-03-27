package casosusos_sincronizacion_productos

import (
	"log"
	"ms-sincronizador-tienda/dominio/constantes"
	"ms-sincronizador-tienda/dominio/entidades"
	dominio_repositorios "ms-sincronizador-tienda/dominio/repositorios/db"
	"strconv"
)

type RecuperarPeticion struct {
	InformacionEds  dominio_repositorios.IInformacionEds
	WacherParametro dominio_repositorios.IRecuperarWacher
}

func (RP *RecuperarPeticion) Ejecutar() (*entidades.HttpRequest, error) {
	info, err := RP.InformacionEds.Ejecutar()
	if err != nil {
		log.Println(constantes.Red + "No se encontro informacion eds" + constantes.Reset)
		return nil, err
	}

	parametro, err := RP.WacherParametro.Consultar("HOST_SERVER")
	if err != nil {
		log.Println(constantes.Red + "No se encontro parametro HOST_SERVER" + constantes.Reset)
		return nil, err
	}

	peticion := &entidades.HttpRequest{
		Metodo: "GET",
		Url:    parametro.Valor + constantes.MS_CLOUD_SINCRONIZADOR + constantes.API_PRODUCTOS_TIENDA + "/" + strconv.FormatInt(info.EdsId, 10),
	}

	return peticion, nil
}
