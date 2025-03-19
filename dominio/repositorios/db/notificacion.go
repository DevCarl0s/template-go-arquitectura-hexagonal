package dominio_repositorios

import "ms-sincronizador-tienda/dominio/entidades"

type INotificacion interface {
	ObtenerPendientes() ([]*entidades.Notificacion, error)
	MarcarProcesada(id int) error
}
