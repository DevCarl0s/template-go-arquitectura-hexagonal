package main

import (
	"log"
	"ms-sincronizador-tienda/presentacion/contenedor"
	"os"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lmicroseconds)
	log.SetOutput(os.Stdout)
	log.Println(`MS-PROCESADOR-NOTIFICACIONES`)

	if err := contenedor.InicializarContenedor(); err != nil {
		log.Fatal(err)
		panic(err)
	}

	go contenedor.SupervisarNotificaciones.Iniciar()
	contenedor.GestionarNotificaciones.Observar()
}
