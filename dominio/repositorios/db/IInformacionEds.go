package dominio_repositorios

import "ms-sincronizador-tienda/dominio/entidades"

type IInformacionEds interface {
	Ejecutar() (*entidades.ParametrosEds, error)
}
