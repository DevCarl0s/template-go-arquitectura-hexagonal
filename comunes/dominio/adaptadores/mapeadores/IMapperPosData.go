package mapeadores

import "ms-sincronizador-tienda/dominio/entidades"

type MapearDatosPos interface {
	MapearA(pos []interface{}) (*entidades.PosData, error)
}
