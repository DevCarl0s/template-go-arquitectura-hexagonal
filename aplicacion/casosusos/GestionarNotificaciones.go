package casosusos

import (
	"log"
	"ms-sincronizador-tienda/dominio/constantes"
	"ms-sincronizador-tienda/dominio/entidades"
	dominio_notificacion "ms-sincronizador-tienda/dominio/notificacion"
	dominio_repositorios "ms-sincronizador-tienda/dominio/repositorios/db"
	"time"
)

type GestionarNotificaciones struct {
	CanalEventos       chan *entidades.Notificacion
	GestorObservadores *dominio_notificacion.GestorObservadores
	Notificaciones     dominio_repositorios.INotificacion
}

func (GN *GestionarNotificaciones) Observar() {

	for {
		log.Println(constantes.Cyan + "Esperando notificaciones" + constantes.Reset)
		notificacion, abierto := <-GN.CanalEventos
		log.Println(constantes.Green + "Notificacion captada" + constantes.Reset)

		if !abierto {
			log.Println(constantes.Red + "CANAL FUE CERRADO, REINICIAR EL SERVICIO" + constantes.Reset)
			time.Sleep(1 * time.Second)
			return
		}

		notificacionObservador := dominio_notificacion.Notificacion{
			ID:   notificacion.ID,
			Tipo: dominio_notificacion.TipoNotificacion(notificacion.TipoNotificacion),
			Data: notificacion.Data,
		}

		if err := GN.GestorObservadores.NotificarObservadores(notificacionObservador); err != nil {
			log.Printf("Error al procesar notificación %d: %v", notificacion.ID, err)
			GN.Notificaciones.MarcarProcesada(notificacion.ID)
			return
		}

		log.Printf("Notificación %d procesada exitosamente", notificacion.ID)
		GN.Notificaciones.MarcarProcesada(notificacion.ID)

	}

}
