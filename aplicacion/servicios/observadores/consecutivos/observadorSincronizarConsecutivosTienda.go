package servicios_observadores_consecutivos

import (
	"fmt"
	dominio_notificacion "ms-sincronizador-tienda/dominio/notificacion"
)

type ObservadorSincronizarConsecutivosTienda struct {
}

func (OP *ObservadorSincronizarConsecutivosTienda) ProcesarNotificacion(notificacion dominio_notificacion.Notificacion) error {
	fmt.Printf("Procesando notificación de consecutivos: %+v\n", notificacion)

	return nil
}

func (o *ObservadorSincronizarConsecutivosTienda) ObtenerTipo() dominio_notificacion.TipoNotificacion {
	return dominio_notificacion.TIPO_CONSECUTIVOS
}
