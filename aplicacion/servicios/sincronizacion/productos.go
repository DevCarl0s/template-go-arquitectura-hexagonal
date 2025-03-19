package servicios_sincronizacion

import (
	"log"
	casosusos_sincronizacion "ms-sincronizador-tienda/aplicacion/casosusos/sincronizacion"
	casosusos_sincronizacion_productos "ms-sincronizador-tienda/aplicacion/casosusos/sincronizacion/productos"
	"ms-sincronizador-tienda/dominio/constantes"
	entidades_sincronizacion_productos "ms-sincronizador-tienda/dominio/entidades/sincronizacion/productos"
)

type Productos struct {
	RecuperarPeticionCloud *casosusos_sincronizacion_productos.RecuperarPeticion
	ConsultarProductos     *casosusos_sincronizacion_productos.ConsultarProductos
	ProcesarInformacion    *casosusos_sincronizacion.ProcesarInformacion
}

func (P *Productos) Ejecutar() (*entidades_sincronizacion_productos.SalidaProductos, error) {
	peticion, err := P.RecuperarPeticionCloud.Ejecutar()
	if err != nil {
		return nil, err
	}

	productos, err := P.ConsultarProductos.Ejecutar(peticion)
	if err != nil {
		log.Println("No se pudieron obtener los productos Cloud")
		return nil, err
	}

	P.ProcesarInformacion.Ejecutar(productos, constantes.SINCRONIZAR_PRODUCTOS)

	return nil, nil
}
