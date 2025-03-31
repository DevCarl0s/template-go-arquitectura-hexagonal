package servicios

import (
	"log"
	"ms-sincronizador-tienda/dominio/constantes"
	"ms-sincronizador-tienda/dominio/entidades"
	dominio_repositorios "ms-sincronizador-tienda/dominio/repositorios/db"
	"time"
)

type SupervisarNotificaciones struct {
	Notificaciones      dominio_repositorios.INotificacion
	CanalNotificaciones chan *entidades.Notificacion
}

func (SN *SupervisarNotificaciones) Iniciar() {
	for {
		notificaciones, err := SN.Notificaciones.ObtenerPendientes()

		if err != nil {
			log.Println(constantes.Red+"Error al consultar procesos", err.Error()+constantes.Reset)
		} else {
			if len(notificaciones) == 0 {
				log.Println(constantes.Yellow + "No hay notificaciones pendientes" + constantes.Reset)
				time.Sleep(time.Second * 10)
			}

			for _, notificacion := range notificaciones {
				SN.CanalNotificaciones <- notificacion
			}
		}
		time.Sleep(time.Second * 1)
	}
}
