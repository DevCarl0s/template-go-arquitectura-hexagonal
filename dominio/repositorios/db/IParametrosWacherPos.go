package dominio_repositorios

import "ms-sincronizador-tienda/dominio/entidades"

type IRecuperarWacher interface {
	Consultar(codigo string) (*entidades.ParametrosWatcher, error)
}
