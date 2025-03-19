package mapper_infraestructura

import (
	"errors"
	"log"
	"ms-sincronizador-tienda/dominio/entidades"
)

type MapaerPosData struct {
}

func (MPD *MapaerPosData) MapearA(pos []interface{}) (*entidades.PosData, error) {

	posDatos := &entidades.PosData{}
	equipoId, ok := pos[0].(int64)
	if ok {
		posDatos.EquiposId = equipoId
	} else {
		return nil, errors.New("campo equipoID  Invalido ")
	}

	host, ok := pos[1].(string)
	if ok {
		posDatos.Host = host
	} else {
		return nil, errors.New("campo host  Invalido ")
	}

	isla, ok := pos[2].(int64)
	if ok {
		posDatos.Isla = isla
	} else {
		return nil, errors.New("campo isla  Invalido ")
	}

	esPosPrincipal, ok := pos[4].(bool)
	if ok {
		posDatos.CheckPosPrincipal = esPosPrincipal
	} else {
		log.Println("campo CheckPosPrincipal  Invalido ")
		//return nil, errors.New("campo CheckPosPrincipal  Invalido ")
	}
	log.Println("POS SECUNDARIO", posDatos)
	return posDatos, nil
}
