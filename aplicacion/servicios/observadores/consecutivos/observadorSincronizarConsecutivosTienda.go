package servicios_observadores_consecutivos

import (
	"fmt"
	"log"
	casosusos_sincronizacion "ms-sincronizador-tienda/aplicacion/casosusos/sincronizacion"
	casosusos_sincronizacion_consecutivos "ms-sincronizador-tienda/aplicacion/casosusos/sincronizacion/consecutivos"
	"ms-sincronizador-tienda/dominio/constantes"
	dominio_notificacion "ms-sincronizador-tienda/dominio/notificacion"
)

type ObservadorSincronizarConsecutivosTienda struct {
	RecuperarPeticionCloud *casosusos_sincronizacion_consecutivos.RecuperarPeticion
	ConsultarConsecutivos  *casosusos_sincronizacion_consecutivos.ConsultarConsecutivos
	ProcesarInformacion    *casosusos_sincronizacion.ProcesarInformacion
}

func (OSCT *ObservadorSincronizarConsecutivosTienda) ProcesarNotificacion(notificacion dominio_notificacion.Notificacion) error {
	fmt.Printf("Procesando notificación de consecutivos: %+v\n", notificacion)

	peticion, err := OSCT.RecuperarPeticionCloud.Ejecutar()
	if err != nil {
		return err
	}
	log.Println("peticion: ", peticion)

	consecutivos, err := OSCT.ConsultarConsecutivos.Ejecutar(peticion)
	if err != nil {
		log.Println("No se pudieron obtener los productos Cloud")
		return err
	}

	err = OSCT.ProcesarInformacion.Ejecutar(consecutivos.Datos, constantes.PROCESAR_RESOLUCIONES)
	if err != nil {
		return err
	}
	return nil
}

func (o *ObservadorSincronizarConsecutivosTienda) ObtenerTipo() dominio_notificacion.TipoNotificacion {
	return dominio_notificacion.TIPO_CONSECUTIVOS
}
